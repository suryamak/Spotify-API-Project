package app

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/suryamak/Spotify-API-Project/pkg/auth"
	"github.com/suryamak/Spotify-API-Project/pkg/spotify_objects"
	"github.com/pkg/errors"
)

type App struct {
	router *gin.Engine
	auth   *auth.Auth
}

func (a *App) Init() {
	a.auth = &auth.Auth{}
	a.router = gin.Default()
	a.router.GET("/helloworld", helloworld)
	a.router.GET("/", func(c *gin.Context) {
		err := a.auth.ExtractAppRegistration()
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect(http.StatusTemporaryRedirect, "https://accounts.spotify.com/authorize"+"?client_id="+a.auth.GetClientID()+"&response_type=code"+"&scope=user-read-private%20user-read-email%20user-top-read"+"&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback"+"&state=quickbrownfox")
	})
	a.router.GET("/callback", func(c *gin.Context) {
		authCode, ok := c.GetQuery("code")
		if ok {
			resp, err := a.auth.RequestToken(authCode)
			if err != nil {
				fmt.Println(err)
				c.Data(http.StatusOK, "text/html; charset=utf-8", resp)
				return
			} else {

				client := &http.Client{}
				req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/artists", nil)
				req.Header.Add("Authorization", "Bearer " + a.auth.Token.AccessToken)
				resp, err := client.Do(req)
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				pagingObject := &spotify_objects.PagingObject{}
				err = json.Unmarshal(body, pagingObject)
				if err != nil {
					fmt.Println(errors.Wrap(err, "[OOF-AUTH] Error unmarshalling paging object"))
				}
				c.JSON(http.StatusOK, pagingObject)
			}
		} else {
			c.String(http.StatusNotFound, "Could not get authorization code")
		}
	})

}

func (a *App) Run() {
	a.router.Run() // localhost:8080
}

func helloworld(c *gin.Context) {
	c.String(200, "Hello World")
}
