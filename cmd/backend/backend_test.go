package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const AUTHCOOKIE_VALUE = "jwt"

func TestBackendNoCookie(t *testing.T) {
	app := &app{randomClient: &mocks.RandomServiceClient{}, characterClient: &mocks.CharacterServiceClient{},
		authClient: &mocks.AuthServiceClient{}}
	testHandler, rr, req := prepareSUT(t, app)
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBackendErrorAuthCookie(t *testing.T) {
	app := &app{randomClient: &mocks.RandomServiceClient{}, characterClient: &mocks.CharacterServiceClient{},
		authClient: errorAuthCookieClient()}

	testHandler, rr, req := prepareSUT(t, app)
	req.AddCookie(&http.Cookie{Name: COOKIE_JWT_NAME, Value: AUTHCOOKIE_VALUE})
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBackend(t *testing.T) {

	tests := []struct {
		randomClient    *mocks.RandomServiceClient
		characterClient *mocks.CharacterServiceClient
		statusCode      int
	}{
		{errorRandomClient(), nil, 500},
		{validRandomClient(), errorCharacterClient(), 500},
		{validRandomClient(), noCharacterClient(), 404},
		{validRandomClient(), validCharacterClient(), 200},
	}

	assert := assert.New(t)
	for _, tt := range tests {

		app := &app{randomClient: tt.randomClient, characterClient: tt.characterClient,
			authClient: validCookieClient()}
		testHandler, rr, req := prepareSUT(t, app)
		req.AddCookie(&http.Cookie{Name: COOKIE_JWT_NAME, Value: AUTHCOOKIE_VALUE})
		testHandler.ServeHTTP(rr, req)

		assert.Equal(tt.statusCode, rr.Code, "handler returned wrong status code: got %v want %v",
			rr.Code, tt.statusCode)
	}

}

func prepareSUT(t *testing.T, app *app) (http.Handler, *httptest.ResponseRecorder, *http.Request) {
	handler := http.HandlerFunc(app.ServeHTTP)
	testHandler := corsMiddleware(app.securedMiddleware((logMiddleware(handler))))

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/random", nil)
	if err != nil {
		t.Fatal(err)
	}

	return testHandler, rr, req

}

func errorAuthCookieClient() *mocks.AuthServiceClient {
	authMockClient := &mocks.AuthServiceClient{}
	authMockClient.On("GetUser", mock.Anything, &auth.Token{Value: AUTHCOOKIE_VALUE}).Return(nil, fmt.Errorf("error"))
	return authMockClient
}

func validCookieClient() *mocks.AuthServiceClient {
	authMockClient := &mocks.AuthServiceClient{}
	authMockClient.On("GetUser", mock.Anything, &auth.Token{Value: AUTHCOOKIE_VALUE}).Return(&auth.User{}, nil)
	return authMockClient
}

func errorRandomClient() *mocks.RandomServiceClient {
	randomClient := &mocks.RandomServiceClient{}
	randomClient.On("GetRandom", mock.Anything, &random.RandomInput{Max: 1000}).Return(nil, fmt.Errorf("error"))
	return randomClient
}

func validRandomClient() *mocks.RandomServiceClient {
	randomClient := &mocks.RandomServiceClient{}
	randomClient.On("GetRandom", mock.Anything, &random.RandomInput{Max: 1000}).Return(&random.RandomNumber{Number: 10}, nil)
	return randomClient
}

func errorCharacterClient() *mocks.CharacterServiceClient {
	characterClient := &mocks.CharacterServiceClient{}
	characterClient.On("GetCharacter", mock.Anything, &character.Input{Number: "10"}).Return(nil, fmt.Errorf("error"))
	return characterClient
}

func noCharacterClient() *mocks.CharacterServiceClient {
	characterClient := &mocks.CharacterServiceClient{}
	characterClient.On("GetCharacter", mock.Anything, &character.Input{Number: "10"}).Return(&character.Output{}, nil)
	return characterClient
}

func validCharacterClient() *mocks.CharacterServiceClient {
	characterClient := &mocks.CharacterServiceClient{}
	characterClient.On("GetCharacter", mock.Anything, &character.Input{Number: "10"}).Return(&character.Output{Id: 10, Name: "name"}, nil)
	return characterClient
}
