package post

import (
	dom "pokeRecords/pkg/domain/records"
	"pokeRecords/pkg/usesCases/handlers/records"

	"github.com/labstack/echo/v4"
)

func CreateRecord(c echo.Context) error {
	s := dom.CreateRecordRequest{}
	if err := c.Bind(&s); err != nil {
		return err
	}
	r := records.Handler{}
	r.NewRecordsHandler()
	response := r.Post.CreateRecord(s)
	return c.JSON(response.StatusCode, response.RecordID)
}
