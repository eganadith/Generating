package main

import (
	"Generating/pkg/apis"
	"github.com/labstack/echo/v4"
)

func main() {
	//echo connection
	e := echo.New()

	apis.EchoConnection(e)

	e.Logger.Fatal(e.Start(":8000"))

}
