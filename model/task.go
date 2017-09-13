package model

import "time"
import "fmt"

// Task represent a ligne in task Table
type Task struct {
	ID          int       `json:"id"`
	Creation    time.Time `json:"creation"`
	Description string    `json:"description"`
}

// Tasks represent a collection of tasks
type Tasks []Task

func (task Task) String() string {
	return fmt.Sprintf("Task %d : '%v' (créée le %v)\n", task.ID, task.Description, task.Creation.Format("02/01/2006"))
}
