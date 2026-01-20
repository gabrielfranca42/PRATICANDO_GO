package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabrielfranca42/simple-go-mod/models"
	"github.com/gorilla/mux"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (taskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, request *http.Request) {
	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)

	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}

func (taskHandler *TaskHandler) CreateTask(write http.ResponseWriter, request *http.Request) {
	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}
}

func (taskHandler *TaskHandler) DeleteTask(write http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(write, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return

	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(write, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (taskHandler *TaskHandler) UpdateTask(write http.ResponseWriter, request *http.Request) {
	_, err := taskHandler.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4",
	)

	if err != nil {
		http.Error(write, err.Error(), http.StatusBadRequest)
		return
	}
}
