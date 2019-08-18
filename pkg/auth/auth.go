package auth

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Auth struct {
	clientID     string
	clientSecret string
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

func (a *Auth) RequestAuthorization() ([]byte, error) {

	resp, err := http.Get("https://accounts.spotify.com/authorize")
	if err != nil {
		return nil, errors.Wrap(err, "[OOF-AUTH] Error sending request to Spotify Accounts Service")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "[OOF-AUTH] Error reading response from Spotify Accounts Service")
	}
	return body, nil

}
