package records

import (
	"pokeRecords/infrastructure/entryPoint/server_http/routes/records/post"
	"pokeRecords/pkg/usesCases/handlers/records"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	handlers := records.Handler{}
	handlers.NewRecordsHandler()
	post.InitPost(g, handlers)

}
