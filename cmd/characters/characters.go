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
	ApiClient ApiClient
	Cache     *cache.Cache
}

type ApiClient struct {
	baseURL string
	Client  http.Client
}

func (a *App) GetCharacter(ctx context.Context, input *character.Input) (*character.Output, error) {
	log.Printf("received call for input %s", input.Number)

	if x, found := a.Cache.Get(input.Number); found {
		log.Println("Getting from cache...")
		return x.(*character.Output), nil
	}

	path := fmt.Sprintf("%s%s", a.ApiClient.baseURL, input.Number)
	resp, err := a.ApiClient.Client.Get(path)
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

	log.Println(character)

	a.Cache.Set(input.Number, character, cache.NoExpiration)

	return character, nil
}

func main() {

	apiClient := ApiClient{
		baseURL: "https://rickandmortyapi.com/api/character/",
		Client: http.Client{
			Timeout: time.Second * 10,
		},
	}
	c := cache.New(cache.NoExpiration, 10*time.Minute)

	app := &App{ApiClient: apiClient, Cache: c}

	grpcServer := server.NewGRPC()
	character.RegisterCharacterServiceServer(grpcServer, app)
	if err := server.ServeGRPC(grpcServer); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
