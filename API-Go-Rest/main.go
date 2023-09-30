package main

import (
	"fmt"

	"github.com/AxelGeorg/API-Go-Rest/database"
	"github.com/AxelGeorg/API-Go-Rest/routes"
)

func main() {
	database.ConectaBancoDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
