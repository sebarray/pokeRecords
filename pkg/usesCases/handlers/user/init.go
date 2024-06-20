package user

import (
	"pokeRecords/pkg/usesCases/handlers/records/post"
	"pokeRecords/pkg/usesCases/handlers/user/delete"
	"pokeRecords/pkg/usesCases/handlers/user/get"
	"pokeRecords/pkg/usesCases/handlers/user/put"
)

type UserHandler struct {
	Get    get.Iget
	Post   post.Ipost
	Put    put.Iput
	Delete delete.Idelete
}

func (h *UserHandler) NewUserHandler() {
	h.Get = &get.Get{}
	h.Post = &post.Post{}
	h.Put = &put.Put{}
	h.Delete = &delete.Delete{}
}
