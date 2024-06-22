package post

import (
	"pokeRecords/pkg/domain/records"
)

func (p Post) CreateRecord(Record records.CreateRecordRequest) records.CreateRecordResponse {
	p.RepositoryRecords.NewRecordsRepository()
	response := p.RepositoryRecords.Post.CreateRecord(Record)
	return response
}
