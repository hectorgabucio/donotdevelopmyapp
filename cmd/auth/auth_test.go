package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"gopkg.in/h2non/gock.v1"
)

const STATE_VALUE = "state"

func TestBackendCallback(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution
	gock.New(AUTH_GOOGLE_URL + "accessToken").
		Get("/").
		Reply(200).
		JSON(map[string]string{"id": "userid"})

	encodedState := base64.URLEncoding.EncodeToString([]byte(STATE_VALUE))

	tests := []struct {
		setCookie   bool
		cookieValue string
		canDecrypt  bool
		formState   string
		formCode    string
		statusCode  int
	}{
		{false, "", true, "formState", "formCode", http.StatusTemporaryRedirect},
		{true, "abc", true, "formState", "formCode", http.StatusTemporaryRedirect},
		{true, "YQ==", false, "formState", "formCode", http.StatusTemporaryRedirect},
		{true, "YQ==", true, "formState", "formCode", http.StatusTemporaryRedirect},
		{true, "YQ==", true, encodedState, "errorCode", http.StatusTemporaryRedirect},
		{true, "YQ==", true, encodedState, "formCode", http.StatusTemporaryRedirect},
	}

	assert := assert.New(t)
	for _, tt := range tests {
		mockConfig := &mocks.ConfigProvider{}
		mockConfig.On("Get", "STATE_SECRET").Return("thisisnotproductionlulz111111111")
		mockOAuth := &mocks.OAuthProvider{}
		mockOAuth.On("Exchange", mock.Anything, "formCode").Return(&oauth2.Token{AccessToken: "accessToken"}, nil)
		mockOAuth.On("Exchange", mock.Anything, "errorCode").Return(nil, fmt.Errorf("error code"))
		mockCipher := &mocks.Cipher{}

		var errDecrypt error
		if !tt.canDecrypt {
			errDecrypt = fmt.Errorf("error decrypt")
		}

		mockCipher.On("Decrypt", []byte("thisisnotproductionlulz111111111"), mock.Anything).Return([]byte(STATE_VALUE), errDecrypt)

		server := &myAuthServiceServer{config: mockConfig, oauth2Config: mockOAuth, cipherUtil: mockCipher}
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
	}

}

func TestGoogleLogin(t *testing.T) {

	mockConfig := &mocks.ConfigProvider{}
	mockConfig.On("Get", "STATE_SECRET").Return("thisisnotproductionlulz111111111")

	mockOAuth := &mocks.OAuthProvider{}
	mockOAuth.On("AuthCodeURL", mock.AnythingOfType("string")).Return("URLRedirect")
	server := &myAuthServiceServer{config: mockConfig, oauth2Config: mockOAuth}

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
