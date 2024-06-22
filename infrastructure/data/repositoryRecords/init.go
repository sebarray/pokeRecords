package records

import (
	"pokeRecords/infrastructure/data/repositoryRecords/delete"
	"pokeRecords/infrastructure/data/repositoryRecords/get"
	"pokeRecords/infrastructure/data/repositoryRecords/post"
	"pokeRecords/infrastructure/data/repositoryRecords/put"
)

type Repository struct {
	Get    get.Iget
	Post   post.Ipost
	Put    put.Iput
	Delete delete.Idelete
}

func (h *Repository) NewRecordsRepository() {
	h.Get = &get.Get{}
	h.Post = &post.Post{}
	h.Put = &put.Put{}
	h.Delete = &delete.Delete{}
}
