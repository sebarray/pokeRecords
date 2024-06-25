package post

import (
	"pokeRecords/infrastructure/data"
	"pokeRecords/pkg/domain/records"
)

func (Post) CreateRecord(Record records.CreateRecordRequest) (records.CreateRecordResponse, error) {
	db := data.GetConnection()
	defer db.Close()
	db.NamedExec("INSERT INTO pokerecord.RECORD ('COLOR', 'VALUE', 'NAME') VALUES (:COLOR, :VALUE, );", Record)
	return records.CreateRecordResponse{}, nil
}
