package post

import (
	"log"
	"pokeRecords/pkg/domain/records"
)

func (p Post) CreateRecord(Record records.CreateRecordRequest) records.CreateRecordResponse {
	p.RepositoryRecords.NewRecordsRepository()
	response, err := p.RepositoryRecords.Post.CreateRecord(Record)
	if err != nil {
		log.Println(err.Error())

	}
	return response
}
