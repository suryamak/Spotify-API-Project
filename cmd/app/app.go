package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suryamak/Spotify-API-Project/pkg/auth"
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
		a.auth.ExtractAppRegistration()
		c.Redirect(http.StatusTemporaryRedirect, "https://accounts.spotify.com/authorize"+"?client_id="+a.auth.GetClientID()+"&response_type=code"+"&scope=user-read-private%20user-read-email"+"&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fcallback"+"&state=quickbrownfox")
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
				c.JSON(http.StatusOK, a.auth.Token)
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
