package post

import (
	"pokeRecords/pkg/domain/records"

	"github.com/jmoiron/sqlx"
)

type Ipost interface {
	CreateRecord(Record records.CreateRecordRequest, db *sqlx.DB) (records.CreateRecordResponse, error)
}

type Post struct {
}
