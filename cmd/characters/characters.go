package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/data"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"
)

type App struct {
	Api         ApiController
	CacheClient data.CacheClient
}

type ApiController interface {
	GetBaseUrl() string
	Get(path string) (*http.Response, error)
}

type ApiClient struct {
	baseURL string
	client  http.Client
}

func (api *ApiClient) GetBaseUrl() string {
	return api.baseURL
}

func (api *ApiClient) Get(path string) (*http.Response, error) {
	return api.client.Get(path)
}

func (a *App) GetCharacter(ctx context.Context, input *character.Input) (*character.Output, error) {
	log.Printf("received call for input %s", input.Number)

	var characterCache *character.Output
	if err := a.CacheClient.Get(input.Number, characterCache); err == nil {
		log.Println("Getting from cache...")
		return characterCache, nil
	}

	path := fmt.Sprintf("%s%s", a.Api.GetBaseUrl(), input.Number)
	resp, err := a.Api.Get(path)
	if err != nil {
		log.Printf("Error while getting characters: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	character := &character.Output{}

	err = json.NewDecoder(resp.Body).Decode(character)
	if err != nil {
		log.Printf("Error decoding response: %s", err)
		return nil, err
	}

	if character.Name != "" {
		if err := a.CacheClient.Set(input.Number, character, time.Hour*24); err != nil {
			log.Println("Error while saving to cache: ", err.Error())
		}
	}

	return character, nil
}

func main() {

	apiClient := ApiClient{
		baseURL: "https://rickandmortyapi.com/api/character/",
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}

	c := data.NewCacheClient()

	app := &App{Api: &apiClient, CacheClient: c}

	grpcServer := server.NewGRPC()
	character.RegisterCharacterServiceServer(grpcServer, app)
	if err := server.ServeGRPC(grpcServer); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
