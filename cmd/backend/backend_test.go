package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/auth"
	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/stretchr/testify/mock"
)

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

	authMockClient := &mocks.AuthServiceClient{}
	authMockClient.On("GetUser", mock.Anything, &auth.Token{Value: "jwt"}).Return(nil, fmt.Errorf("error"))
	app := &app{randomClient: &mocks.RandomServiceClient{}, characterClient: &mocks.CharacterServiceClient{},
		authClient: authMockClient}

	testHandler, rr, req := prepareSUT(t, app)
	testHandler.ServeHTTP(rr, req)
	req.AddCookie(&http.Cookie{Name: COOKIE_JWT_NAME, Value: "jwt"})
	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
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

/*


func TestGetRandomNumber(t *testing.T) {

	tests := []struct {
		maxInput     int64
		errorMessage string
	}{
		{10, ""},
		{-1, "rpc error: code = Unknown desc = Max input cant be 0 or lesser than 0"},
	}
	init := make(chan bool)
	s := grpctest.InitServer(init)
	defer s.GracefulStop()
	random.RegisterRandomServiceServer(s, &randomHandler{})
	init <- true

	conn := grpctest.Dialer()
	defer conn.Close()

	assert := assert.New(t)

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%s", tt.maxInput, tt.errorMessage)
		t.Run(testname, func(t *testing.T) {
			ctx := context.Background()
			client := random.NewRandomServiceClient(conn)
			_, err := client.GetRandom(ctx, &random.RandomInput{Max: tt.maxInput})
			if err != nil {
				assert.Equal(tt.errorMessage, err.Error(), "It should be equal")
			}

		})
	}

*/
