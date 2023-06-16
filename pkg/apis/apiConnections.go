package apis

import (
	"github.com/labstack/echo/v4"
)

func EchoConnection(e *echo.Echo) {

	g := e.Group("/userinput")
	CustomerApis(g)

}

func CustomerApis(g *echo.Group) {
	g.POST("/data", createData)
	g.GET("/data", getData)
	g.GET("/trig", triggerAPI)

}
