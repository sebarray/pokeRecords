package post

import (
	"log"
	"pokeRecords/infrastructure/data"
	"pokeRecords/pkg/domain/records"
)

func (p Post) CreateRecord(Record records.CreateRecordRequest) records.CreateRecordResponse {

	db := data.GetConnection()

	p.RepositoryRecords.NewRecordsRepository()
	response, err := p.RepositoryRecords.Post.CreateRecord(Record, db)
	if err != nil {
		log.Println(err.Error())
		return response
	}

	return response
}
