package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rjtokenring/goms/stringstxt"
	"net/http"
)

func main() {
	e := echo.New()

	initGetHandler(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func initGetHandler(e *echo.Echo)  {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! ")
	})
	e.GET("/txt/reverse/:txtr", func(c echo.Context) error {
		var txtToReverse = c.Param("txtr")
		txtToReverse = "This is the reversed txt: " + stringstxt.ReverseRunes(txtToReverse)
		return c.String(http.StatusOK, txtToReverse)
	})
}

