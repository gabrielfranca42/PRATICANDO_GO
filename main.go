package main

import (
	"log"
	"net/http"

	"github.com/gabrielfranca42/simple-go-mod/config"
	"github.com/gabrielfranca42/simple-go-mod/models"
	"github.com/gorilla/mux"
)

func main() {

	db := config.SetupDataBase()
	defer db.Close()

	_, err := db.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"api funcionando"}`))
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
