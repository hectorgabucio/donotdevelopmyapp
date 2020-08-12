package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestBackendNoCookie(t *testing.T) {

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

func prepareSUTGoogleLogin(t *testing.T, server *myAuthServiceServer) (http.Handler, *httptest.ResponseRecorder, *http.Request) {
	handler := http.HandlerFunc(server.handleGoogleLogin)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	return handler, rr, req
}
