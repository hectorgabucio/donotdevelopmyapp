package oauth

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const AUTH_GOOGLE_URL = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleActions interface {
	GetUserDataFromGoogle(code string) ([]byte, error)
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
}

type GoogleActionsImpl struct {
	oauth2Provider OAuthProvider
}

type OAuthProvider interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

func NewGoogleActions() GoogleActions {
	return &GoogleActionsImpl{oauth2Provider: NewGoogleOauth()}
}

func NewGoogleOauth() OAuthProvider {
	return &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func (g *GoogleActionsImpl) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return g.oauth2Provider.AuthCodeURL(state, opts...)
}

func (g *GoogleActionsImpl) GetUserDataFromGoogle(code string) ([]byte, error) {
	token, err := g.oauth2Provider.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(AUTH_GOOGLE_URL + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
