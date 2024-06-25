package server_http

import (
	"pokeRecords/infrastructure/entryPoint/server_http/routes/records"

	"github.com/labstack/echo/v4"
)

type Server struct {
}

func (Server) NewServer() {
	e := echo.New()
	Route(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func Route(e *echo.Echo) {
	records.Init(e.Group("/records"))

}
