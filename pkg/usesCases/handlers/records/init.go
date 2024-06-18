package records

import (
	"pokeRecords/pkg/usesCases/handlers/records/delete"
	"pokeRecords/pkg/usesCases/handlers/records/get"
	"pokeRecords/pkg/usesCases/handlers/records/post"
	"pokeRecords/pkg/usesCases/handlers/records/put"
)

type RecordsHandler struct {
	Get    get.Iget
	Post   post.Ipost
	Put    put.Iput
	Delete delete.Idelete
}

func (h *RecordsHandler) NewRecordsHandler() {
	h.Get = &get.Get{}
	h.Post = &post.Post{}
	h.Put = &put.Put{}
	h.Delete = &delete.Delete{}
}
