package main

import "pokeRecords/infrastructure/entryPoint/server_http"

func init() {

}

func main() {
	s := server_http.Server{}
	s.NewServer()
}
