package post

import "pokeRecords/pkg/domain/records"

type Ipost interface {
	CreateRecord(Record records.CreateRecordRequest) (records.CreateRecordResponse, error)
}

type Post struct {
}
