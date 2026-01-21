package main

import (
	"log"
	"net/http"

	"github.com/gabrielfranca42/simple-go-mod/config"
	"github.com/gabrielfranca42/simple-go-mod/handlers"
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

	taskHandler := handlers.NewTaskHandler(&db)

	router.HandleFunc("tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
