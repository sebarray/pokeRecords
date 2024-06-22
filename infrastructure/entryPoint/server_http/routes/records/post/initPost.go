package post

import "github.com/labstack/echo/v4"

func InitPost(g *echo.Group) {

	g.POST("/records", CreateRecord)

}
