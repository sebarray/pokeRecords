package post

import (
	"pokeRecords/pkg/domain/records"

	"github.com/jmoiron/sqlx"
)

func (Post) CreateRecord(Record records.CreateRecordRequest, db *sqlx.DB) (records.CreateRecordResponse, error) {

	db.Exec("INSERT INTO pokerecord.RECORD ('COLOR', 'VALUE', 'NAME') VALUES (:COLOR, :VALUE, );", Record)
	return records.CreateRecordResponse{}, nil
}
