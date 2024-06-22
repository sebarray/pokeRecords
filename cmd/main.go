package main

import (
	"fmt"
	"pokeRecords/pkg/domain/records"

	"github.com/go-playground/validator"
)

func main() {
	x := records.CreateRecordRequest{
		Email: "test@hotmail..com",
	}
	err := validator.New().Struct(x)
	if err != nil {
		fmt.Println(err.Error())

	}

}
