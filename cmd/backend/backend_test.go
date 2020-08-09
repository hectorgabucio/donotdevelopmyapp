package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
)

func TestBackendNoCookie(t *testing.T) {

	app := &app{randomClient: &mocks.RandomServiceClient{}, characterClient: &mocks.CharacterServiceClient{}, authClient: &mocks.AuthServiceClient{}}
	handler := http.HandlerFunc(app.ServeHTTP)
	testHandler := corsMiddleware(app.securedMiddleware((logMiddleware(handler))))

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/random", nil)
	if err != nil {
		t.Fatal(err)
	}

	testHandler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
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
