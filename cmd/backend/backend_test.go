package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/data"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/random"
	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const AUTHCOOKIE_VALUE = "jwt"

func TestBackendNoCookie(t *testing.T) {
	app := &app{randomClient: &mocks.RandomServiceClient{}, characterClient: &mocks.CharacterServiceClient{},
		authClient: &mocks.AuthServiceClient{}}
	testHandler, rr, req := prepareSUTAddCharacter(t, app)
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBackendErrorAuthCookie(t *testing.T) {
	app := &app{randomClient: &mocks.RandomServiceClient{}, characterClient: &mocks.CharacterServiceClient{},
		authClient: errorAuthCookieClient()}

	testHandler, rr, req := prepareSUTAddCharacter(t, app)
	req.AddCookie(&http.Cookie{Name: COOKIE_JWT_NAME, Value: AUTHCOOKIE_VALUE})
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBackendGetCharacters(t *testing.T) {
	tests := []struct {
		userRepository *mocks.UserRepository
		statusCode     int
	}{
		{mockRepositoryGetCharactersKO(), 500},
		{mockRepositoryGetCharactersEmpty(), 200},
		{mockRepositoryGetCharactersOK(), 200},
	}

	assert := assert.New(t)
	for _, tt := range tests {
		app := &app{userRepository: tt.userRepository, authClient: validCookieClient()}
		testHandler, rr, req := prepareSUTGetCharacters(t, app)
		req.AddCookie(&http.Cookie{Name: COOKIE_JWT_NAME, Value: AUTHCOOKIE_VALUE})
		testHandler.ServeHTTP(rr, req)

		assert.Equal(tt.statusCode, rr.Code, "handler returned wrong status code: got %v want %v",
			rr.Code, tt.statusCode)
	}
}

func TestBackendAddCharacter(t *testing.T) {

	tests := []struct {
		randomClient    *mocks.RandomServiceClient
		characterClient *mocks.CharacterServiceClient
		userRepository  *mocks.UserRepository
		statusCode      int
	}{
		{errorRandomClient(), nil, mockRepositoryAddCharacter(), 500},
		{validRandomClient(), errorCharacterClient(), mockRepositoryAddCharacter(), 500},
		{validRandomClient(), noCharacterClient(), mockRepositoryAddCharacter(), 404},
		{validRandomClient(), validCharacterClient(), mockRepositoryError(), 500},
		{validRandomClient(), validCharacterClient(), mockRepositoryAddCharacter(), 200},
	}

	assert := assert.New(t)
	for _, tt := range tests {

		app := &app{randomClient: tt.randomClient, characterClient: tt.characterClient,
			authClient: validCookieClient(), userRepository: tt.userRepository}
		testHandler, rr, req := prepareSUTAddCharacter(t, app)
		req.AddCookie(&http.Cookie{Name: COOKIE_JWT_NAME, Value: AUTHCOOKIE_VALUE})
		testHandler.ServeHTTP(rr, req)

		assert.Equal(tt.statusCode, rr.Code, "handler returned wrong status code: got %v want %v",
			rr.Code, tt.statusCode)
	}

}

func prepareSUTAddCharacter(t *testing.T, app *app) (http.Handler, *httptest.ResponseRecorder, *http.Request) {
	handler := http.HandlerFunc(app.AddNewCharacterForUser)
	testHandler := corsMiddleware(app.securedMiddleware((logMiddleware(handler))))

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/characters", nil)
	if err != nil {
		t.Fatal(err)
	}

	return testHandler, rr, req

}

func prepareSUTGetCharacters(t *testing.T, app *app) (http.Handler, *httptest.ResponseRecorder, *http.Request) {
	handler := http.HandlerFunc(app.GetCharactersOfUser)
	testHandler := corsMiddleware(app.securedMiddleware((logMiddleware(handler))))

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/characters/me", nil)
	if err != nil {
		t.Fatal(err)
	}

	return testHandler, rr, req

}

func mockRepositoryError() *mocks.UserRepository {
	userRepository := &mocks.UserRepository{}
	userRepository.On("AddCharacterToUser", &data.Character{ID: "10", Name: "name", Image: ""}, "10").Return(errors.New("error repo"))
	return userRepository
}

func mockRepositoryAddCharacter() *mocks.UserRepository {
	userRepository := &mocks.UserRepository{}
	userRepository.On("AddCharacterToUser", &data.Character{ID: "10", Name: "name", Image: ""}, "10").Return(nil)
	return userRepository
}

func mockRepositoryGetCharactersKO() *mocks.UserRepository {
	userRepository := &mocks.UserRepository{}
	userRepository.On("GetCharactersByUserId", "10").Return(nil, errors.New("error"))
	return userRepository
}

func mockRepositoryGetCharactersEmpty() *mocks.UserRepository {
	userRepository := &mocks.UserRepository{}
	userRepository.On("GetCharactersByUserId", "10").Return([]data.UserCharacter{}, nil)
	return userRepository
}

func mockRepositoryGetCharactersOK() *mocks.UserRepository {
	userRepository := &mocks.UserRepository{}
	userRepository.On("GetCharactersByUserId", "10").Return([]data.UserCharacter{
		{Count: 1, Character: data.Character{ID: "1", Image: "image", Name: "name"}},
		{Count: 2, Character: data.Character{ID: "2", Image: "image", Name: "name"}},
	}, nil)
	return userRepository
}

func errorAuthCookieClient() *mocks.AuthServiceClient {
	authMockClient := &mocks.AuthServiceClient{}
	authMockClient.On("GetUser", mock.Anything, &auth.Token{Value: AUTHCOOKIE_VALUE}).Return(nil, fmt.Errorf("error"))
	return authMockClient
}

func validCookieClient() *mocks.AuthServiceClient {
	authMockClient := &mocks.AuthServiceClient{}
	authMockClient.On("GetUser", mock.Anything, &auth.Token{Value: AUTHCOOKIE_VALUE}).Return(&auth.User{Id: "10"}, nil)
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
