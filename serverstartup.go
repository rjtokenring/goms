package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/rjtokenring/goms/dbaccess"
	"github.com/rjtokenring/goms/serverstub"
	"github.com/rjtokenring/goms/stringstxt"
	"net/http"
)

var serverBinding = ":1323"
var version = "1.0"

func main() {
	//Init DB before starting web server
	var dblink = dbaccess.InitDb()

	e := echo.New()

	initGetHandler(e, dblink)

	e.Logger.Fatal(e.Start(serverBinding))
}

func initGetHandler(e *echo.Echo, db *sql.DB)  {
	//Base
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Go MS test - version: "+version)
	})
	e.GET("/txt/reverse/:txtr", func(c echo.Context) error {
		var txtToReverse = c.Param("txtr")
		txtToReverse = "This is the reversed txt: " + stringstxt.ReverseRunes(txtToReverse)
		return c.String(http.StatusOK, txtToReverse)
	})

	//Server genearted stubs with OpenApi spec
	var implementingStubs = serverstub.GoMsServerImpl{DbLink: db}
	serverstub.RegisterHandlers(e, &implementingStubs)
}

