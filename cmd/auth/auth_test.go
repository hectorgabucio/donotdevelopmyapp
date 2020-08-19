package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const STATE_VALUE = "state"

func TestBackendCallback(t *testing.T) {

	encodedState := base64.URLEncoding.EncodeToString([]byte(STATE_VALUE))

	tests := []struct {
		setCookie   bool
		cookieValue string
		canDecrypt  bool
		formState   string
		formCode    string
		statusCode  int
		jwtSet      bool
	}{
		{false, "", true, "formState", "formCode", http.StatusTemporaryRedirect, false},
		{true, "abc", true, "formState", "formCode", http.StatusTemporaryRedirect, false},
		{true, "YQ==", false, "formState", "formCode", http.StatusTemporaryRedirect, false},
		{true, "YQ==", true, "formState", "formCode", http.StatusTemporaryRedirect, false},
		{true, "YQ==", true, encodedState, "errorCode", http.StatusTemporaryRedirect, false},
		{true, "YQ==", true, encodedState, "formCode", http.StatusFound, true},
	}

	assert := assert.New(t)
	for _, tt := range tests {
		mockConfig := &mocks.ConfigProvider{}
		mockConfig.On("Get", "STATE_SECRET").Return("thisisnotproductionlulz111111111")
		mockConfig.On("Get", "ACCESS_SECRET").Return("thisisnotproductionlulz111111111")
		mockConfig.On("Get", "FRONT_URL").Return("/urlToFront")

		mockAuth := &mocks.GoogleActions{}
		mockAuth.On("GetUserDataFromGoogle", "formCode").Return([]byte(`{"id": "userid", "email": "email"}`), nil)
		mockAuth.On("GetUserDataFromGoogle", "errorCode").Return(nil, fmt.Errorf("error code"))

		mockCipher := &mocks.Cipher{}
		mockJwt := &mocks.JwtProvider{}
		mockJwt.On("CreateToken", "userid", mock.Anything, EXPIRES).Return("jwtToken", nil)

		var errDecrypt error
		if !tt.canDecrypt {
			errDecrypt = fmt.Errorf("error decrypt")
		}
		mockCipher.On("Decrypt", []byte("thisisnotproductionlulz111111111"), mock.Anything).Return([]byte(STATE_VALUE), errDecrypt)

		db, sqlmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(err)
		gdb, err := gorm.Open("postgres", db) // open gorm db
		assert.NoError(err)
		defer gdb.Close()

		rows := sqlmock.
			NewRows([]string{"id", "name"}).
			AddRow("userid", "TOBEGENERATED")
		sqlmock.
			ExpectQuery(`SELECT * FROM "users" WHERE ("users"."id" = $1) AND ("users"."name" = $2) ORDER BY "users"."id" ASC LIMIT 1`).
			WillReturnRows(rows)
		sqlmock.ExpectBegin()
		server := &myAuthServiceServer{config: mockConfig, googleAuth: mockAuth, cipherUtil: mockCipher, db: gdb, jwt: mockJwt}

		testHandler, rr, req := prepareSUTGoogleCallback(t, server)

		if tt.setCookie {
			req.AddCookie(&http.Cookie{Name: COOKIE_STATE_NAME, Value: tt.cookieValue})
		}

		req.Form = url.Values{}
		req.Form.Add("state", tt.formState)
		req.Form.Add("code", tt.formCode)

		testHandler.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.statusCode {
			assert.Failf("error", "handler returned wrong status code: got %d want %d",
				status, tt.statusCode)
		}

		if tt.jwtSet {
			cookies := rr.Result().Cookies()
			if len(cookies) > 0 && cookies[0].Name != COOKIE_JWT_NAME {
				assert.Failf("error", "no cookie jwt found")
			}
		}
	}

}

func TestGoogleLogin(t *testing.T) {
	mockCipher := &mocks.Cipher{}
	mockCipher.On("Encrypt", []byte("thisisnotproductionlulz111111111"), mock.Anything).Return([]byte("encripted"), nil)

	mockConfig := &mocks.ConfigProvider{}
	mockConfig.On("Get", "STATE_SECRET").Return("thisisnotproductionlulz111111111")

	mockAuth := &mocks.GoogleActions{}
	mockAuth.On("AuthCodeURL", mock.AnythingOfType("string")).Return("URLRedirect")

	server := &myAuthServiceServer{config: mockConfig, cipherUtil: mockCipher, googleAuth: mockAuth}

	testHandler, rr, req := prepareSUTGoogleLogin(t, server)
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusTemporaryRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(rr.Result().Cookies()) != 1 {
		t.Errorf("error no cookie")
	}

	if cookieState := rr.Result().Cookies()[0]; cookieState.Name != COOKIE_STATE_NAME {
		t.Errorf("error no random state found in response")
	}

}

func prepareSUTGoogleCallback(t *testing.T, server *myAuthServiceServer) (http.Handler, *httptest.ResponseRecorder, *http.Request) {
	handler := http.HandlerFunc(server.oauthGoogleCallback)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/callback", nil)
	if err != nil {
		t.Fatal(err)
	}
	return handler, rr, req
}

func prepareSUTGoogleLogin(t *testing.T, server *myAuthServiceServer) (http.Handler, *httptest.ResponseRecorder, *http.Request) {
	handler := http.HandlerFunc(server.handleGoogleLogin)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	return handler, rr, req
}
