package post

import (
	"pokeRecords/infrastructure/entryPoint/server_http/routes/records/middleware"
	"pokeRecords/pkg/usesCases/handlers/records"

	"github.com/labstack/echo/v4"
)

type PostRoute struct {
	Handler records.Handler
}

func InitPost(g *echo.Group, Handler records.Handler) {
	var p PostRoute
	p.Handler = Handler
	g.POST("/records", p.CreateRecord, middleware.RecoveryMiddleware)

}
