package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/test/grpctest"
	"github.com/hectorgabucio/donotdevelopmyapp/test/mocks"
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
	init := make(chan bool)
	s := grpctest.InitServer(init)
	defer s.GracefulStop()

	apiClient := mocks.ApiController{}

	apiClient.On("GetBaseUrl").Return("/api/")
	apiClient.On("Get", "/api/10").Return(nil, fmt.Errorf("error10"))
	apiClient.On("Get", "/api/11").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader("malformed json"))}, nil)
	apiClient.On("Get", "/api/12").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"id": 1, "name": "Ricky", "image": "image"}`))}, nil)

	cacheClient := &mocks.CacheClient{}
	cacheClient.On("Get", mock.Anything, mock.Anything).Return(errors.New("error cache"))
	cacheClient.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error cache"))

	app := &App{Api: &apiClient, CacheClient: cacheClient}

	character.RegisterCharacterServiceServer(s, app)
	init <- true

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

/*
func TestGetCharacterCached(t *testing.T) {
	init := make(chan bool)
	s := grpctest.InitServer(init)
	defer s.GracefulStop()

	c := cache.New(cache.NoExpiration, 10*time.Minute)

	apiClient := mocks.ApiController{}

	apiClient.On("GetBaseUrl").Return("/api/")
	apiClient.On("Get", "/api/1").Return(&http.Response{
		Body: ioutil.NopCloser(strings.NewReader(`{"id": 1, "name": "Ricky", "image": "image"}`))}, nil)

	app := &App{Api: &apiClient, Cache: c}

	character.RegisterCharacterServiceServer(s, app)
	init <- true

	conn := grpctest.Dialer()
	defer conn.Close()

	assert := assert.New(t)

	ctx := context.Background()
	client := character.NewCharacterServiceClient(conn)
	resp, _ := client.GetCharacter(ctx, &character.Input{Number: "1"})

	assert.Equal(int32(1), resp.Id, "should be equal")
	apiClient.AssertNumberOfCalls(t, "Get", 1)

	// this time is cached, should not do http call
	resp, _ = client.GetCharacter(ctx, &character.Input{Number: "1"})
	assert.Equal(int32(1), resp.Id, "should be equal")
	apiClient.AssertNumberOfCalls(t, "Get", 1)
}

*/
