package web

import (
	"github.com/bgiegel/TodoREST/repo"
)

// TodoApp est instancié une seule fois
// au lancement de l'application et contient
// des variables global tel que les instances des différents repo.
type TodoApp struct {
	TaskRepo *repo.TaskRepository
}
