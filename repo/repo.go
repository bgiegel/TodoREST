package repo

import (
	"database/sql"
	"log"

	"github.com/bgiegel/TodoREST/model"
)

//DB enveloppe la base de donnée
type Repository struct {
	*sql.DB
}

type TaskStore interface {
	AllTasks() ([]*model.Task, error)
}

type TaskRepository struct {
	*Repository
}

// NewDB initialise la cconnexion à la base de donnée
func NewDB(dataSourceName string) *Repository {
	var err error
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panicln(err)
	}

	if db != nil {
		log.Println("Connection to tododb -> OK")
	}

	return &Repository{db}
}

// GetTaskRepository recupère une instance du repo des taches
func GetTaskRepository() *TaskRepository {
	db := NewDB("postgres://todo_admin:password@tododb/tododb?sslmode=disable")
	return &TaskRepository{db}
}
