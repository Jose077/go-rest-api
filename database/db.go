package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringConexao := "host=172.18.0.2 user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringConexao))

	if err != nil {
		log.Panic("Erro ao conectar com o db!", err.Error())
	}

}
