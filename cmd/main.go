package main

import (
	"fmt"

	"github.com/suryamak/Spotify-API-Project/cmd/app"
)

func main() {

	fmt.Println("Hello World")

	theApp := &app.App{}
	theApp.Init()
	theApp.Run()

}
