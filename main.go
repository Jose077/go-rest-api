package main

import (
	"github.com/jose077/go-rest-api/database"
	"github.com/jose077/go-rest-api/models"
	"github.com/jose077/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "teste", Historia: "historia teste"},
		{Id: 2, Nome: "teste 2", Historia: "historia teste 2"},
	}

	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
