package model

import "time"

// Task represent a ligne in task Table
type Task struct {
	ID          int       `json:"id"`
	Creation    time.Time `json:"creation"`
	Description string    `json:"description"`
}

// Tasks represent a collection of tasks
type Tasks []Task
