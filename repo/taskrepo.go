package repo

import (
	"fmt"

	"github.com/bgiegel/TodoREST/model"
)

var currentID int

var tasks model.Tasks

// Give us some seed data
func init() {
	RepoCreateTask(model.Task{Description: "Write presentation"})
	RepoCreateTask(model.Task{Description: "Host meetup"})
}

func RepoFindTask(id int) model.Task {
	for _, t := range tasks {
		if t.ID == id {
			return t
		}
	}
	// return empty Task if not found
	return model.Task{}
}

func RepoAllTasks() model.Tasks {
	return tasks
}

func RepoCreateTask(t model.Task) model.Task {
	currentID++
	t.ID = currentID
	tasks = append(tasks, t)
	return t
}

func RepoDestroyTask(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Task with id of %d to delete", id)
}
