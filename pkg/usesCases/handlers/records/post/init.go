package post

import (
	records "pokeRecords/infrastructure/data/repositoryRecords"
	recordsDomain "pokeRecords/pkg/domain/records"
)

type Ipost interface {
	CreateRecord(Record recordsDomain.CreateRecordRequest) recordsDomain.CreateRecordResponse
}

type Post struct {
	RepositoryRecords records.Repository
}
