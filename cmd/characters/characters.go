package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hectorgabucio/donotdevelopmyapp/internal/character"
	"github.com/hectorgabucio/donotdevelopmyapp/internal/server"
	"github.com/patrickmn/go-cache"
)

type App struct {
	Api   ApiController
	Cache *cache.Cache
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

func (a *App) GetCharacter(ctx context.Context, input *character.Id) (*character.CharacterResponse, error) {
	log.Printf("received call for input %s", input.Number)

	if x, found := a.Cache.Get(input.Number); found {
		log.Println("Getting from cache...")
		return x.(*character.CharacterResponse), nil
	}

	path := fmt.Sprintf("%s%s", a.Api.GetBaseUrl(), input.Number)
	resp, err := a.Api.Get(path)
	if err != nil {
		log.Printf("Error while getting characters: %s", err)
		return nil, err
	}
	defer resp.Body.Close()
	character := &character.CharacterResponse{}

	err = json.NewDecoder(resp.Body).Decode(character)
	if err != nil {
		log.Printf("Error decoding response: %s", err)
		return nil, err
	}

	if character.Name != "" {
		a.Cache.Set(input.Number, character, cache.NoExpiration)
	}

	return character, nil
}

func (a *App) GetAllCharactersOfUser(ctx context.Context, in *character.Id) (*character.CharactersResponse, error) {
	return nil, nil
}

func main() {

	apiClient := ApiClient{
		baseURL: "https://rickandmortyapi.com/api/character/",
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}
	c := cache.New(cache.NoExpiration, 10*time.Minute)

	app := &App{Api: &apiClient, Cache: c}

	grpcServer := server.NewGRPC()
	character.RegisterCharacterServiceServer(grpcServer, app)
	if err := server.ServeGRPC(grpcServer); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
