package post

import (
	dom "pokeRecords/pkg/domain/records"

	"github.com/labstack/echo/v4"
)

func (p *PostRoute) CreateRecord(c echo.Context) error {
	s := dom.CreateRecordRequest{}
	if err := c.Bind(&s); err != nil {
		return err
	}
	response := p.Handler.Post.CreateRecord(s)
	return c.JSON(response.StatusCode, response.RecordID)
}
