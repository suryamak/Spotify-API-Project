package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type Auth struct {
	clientID     string
	clientSecret string
	Token        *Token
}

func (a *Auth) GetClientID() string {

	return a.clientID

}

func (a *Auth) GetClientSecret() string {

	return a.clientSecret

}

func (a *Auth) ExtractAppRegistration() error {

	err := godotenv.Load()
	if err != nil {
		return errors.Wrap(err, "[OOF-AUTH] Error loading .env file")
	}
	var ok bool
	a.clientID, ok = os.LookupEnv("CLIENT_ID")
	if !ok {
		return errors.New("[OOF-AUTH] The client ID has not been set.")
	}
	a.clientSecret, ok = os.LookupEnv("CLIENT_SECRET")
	if !ok {
		return errors.New("[OOF-AUTH] The client secret has not been set.")
	}
	return nil

}

func (a *Auth) RequestToken(code string) ([]byte, error) {

	resp, err := http.PostForm("https://accounts.spotify.com/api/token",
		url.Values{"client_id": {a.clientID}, "client_secret": {a.clientSecret}, "grant_type": {"authorization_code"}, "code": {code}, "redirect_uri": {"http://localhost:8080/callback"}})
	if err != nil {
		return nil, errors.Wrap(err, "[OOF-AUTH] Error sending token request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "[OOF-AUTH] Error reading token response")
	}
	token := &Token{}
	err = json.Unmarshal(body, token)
	if err != nil {
		return nil, errors.Wrap(err, "[OOF-AUTH] Error unmarshalling JSON")
	}
	a.Token = token
	return body, nil

}
