package models

type Task struct {
	ID          int
	Title       string
	Description string
	Status      bool
}

const (
	TableName      = "tasks"
	CreateTableSQL = `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
)
