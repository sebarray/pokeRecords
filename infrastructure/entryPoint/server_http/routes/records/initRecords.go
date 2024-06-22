package records

import "github.com/labstack/echo/v4"

func Init(g *echo.Group) {

	g.Group("/public")

}
