package main

import (
	"log"
	"net/http"

	"github.com/gabrielfranca42/simple-go-mod/config"
)

//funcao principal e por ela e inciadada
//ponto de entrada da aplicação

func main() {

	dbConnection := config.SetupDataBase()

	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
