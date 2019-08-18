package app

import (
	"fmt"

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
		resp, err := a.auth.RequestAuthorization()
		if err != nil {
			c.String(500, "Error")
			fmt.Println(err)
			return
		} else {
			c.String(200, string(resp))
		}

	})
}

func (a *App) Run() {
	a.router.Run() // listen and serve on 0.0.0.0:8080
}

func helloworld(c *gin.Context) {
	c.String(200, "Hello World")
}
