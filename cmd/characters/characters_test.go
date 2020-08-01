package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/test/grpctest"
	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestGetCharacter(t *testing.T) {

	tests := []struct {
		numberInput  string
		errorMessage string
		character    *character.Output
	}{
		{"10", "rpc error: code = Unknown desc = error10", nil},
		{"11", "rpc error: code = Unknown desc = invalid character 'm' looking for beginning of value", nil},
		{"12", "", &character.Output{Id: 1, Name: "Ricky", Image: "image"}},
	}

	s := grpctest.InitServer()
	defer s.GracefulStop()

	c := cache.New(cache.NoExpiration, 10*time.Minute)

	apiClient := mocks.ApiController{}

	apiClient.On("GetBaseUrl").Return("/api/")
	apiClient.On("Get", "/api/10").Return(nil, fmt.Errorf("error10"))
	apiClient.On("Get", "/api/11").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader("malformed json"))}, nil)
	apiClient.On("Get", "/api/12").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"id": 1, "name": "Ricky", "image": "image"}`))}, nil)

	app := &App{Api: &apiClient, Cache: c}

	character.RegisterCharacterServiceServer(s, app)

	conn := grpctest.Dialer()
	defer conn.Close()

	assert := assert.New(t)

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.numberInput, tt.errorMessage)
		t.Run(testname, func(t *testing.T) {
			ctx := context.Background()
			client := character.NewCharacterServiceClient(conn)
			resp, err := client.GetCharacter(ctx, &character.Input{Number: tt.numberInput})
			if err != nil {
				assert.Equal(tt.errorMessage, err.Error(), "It should be equal")
			} else {
				assert.True(proto.Equal(tt.character, resp), "should be equal")
			}
		})
	}

}
