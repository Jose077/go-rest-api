package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jose077/go-rest-api/database"
	"github.com/jose077/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var p []models.Personalidade

	// busca dados no db
	database.DB.Find(&p)

	// retorna dados para o cliente
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	// busca personalidade por id no db
	database.DB.First(&personalidade, id)

	// retorn json para o cliente
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	// Prepara resposta enviada na requisicao
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	// Cria no banco de dados
	database.DB.Create(&novaPersonalidade)

	// Retornar personalidade para o cliente
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeleteUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	// pega id da rota
	vars := mux.Vars(r)
	id := vars["id"]

	// instancia variavel para referenciar entidade do banco e receber o valor
	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)

	// Retorna para o cliente
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	// Pega id da rota
	vars := mux.Vars(r)
	log.Println(vars)

	id := vars["id"]

	// Instancia da entitie
	var personalidade models.Personalidade

	// busca personalidade por id
	database.DB.First(&personalidade, id)

	// Pega valor do body e insere na variavel de personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)

	// Salva alteração no banco
	database.DB.Save(&personalidade)

	// retorna para o cliente
	json.NewEncoder(w).Encode(personalidade)
}
