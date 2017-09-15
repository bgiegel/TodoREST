package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bgiegel/TodoREST/model"
)

var currentID int

func (repo *TaskRepository) FindTask(id int) model.Task {
	task := model.Task{}
	row := repo.QueryRow("SELECT * FROM task WHERE id=$1", id)
	switch err := row.Scan(&task.ID, &task.Creation, &task.Description); err {
	case sql.ErrNoRows:
		log.Printf("No task found for id = %d \n", id)
	case nil:
		log.Printf("Task found : %v", task)
	default:
		panic(err)
	}

	return task
}

func (repo *TaskRepository) AllTasks() model.Tasks {
	tasks := model.Tasks{}
	rows, err := repo.Query("SELECT * FROM task")
	switch err {
	case sql.ErrNoRows:
		log.Printf("No rows found")
	case nil:
		defer rows.Close()
		log.Printf("rows found : %v", rows)
		tasks = buildTasks(rows)
	default:
		panic(err)
	}
	return tasks
}

func buildTasks(rows *sql.Rows) model.Tasks {
	tasks := make([]model.Task, 0)
	for rows.Next() {
		task := model.Task{}
		err := rows.Scan(&task.ID, &task.Creation, &task.Description)
		if err != nil {
			log.Panicln(err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Panicln(err)
	}

	return tasks
}

func (repo *TaskRepository) CreateTask(t model.Task) (taskID int) {
	err := repo.QueryRow("INSERT INTO task(description) VALUES($1) RETURNING id", t.Description).
		Scan(&taskID)
	if err != nil {
		log.Panic(err)
	}
	return
}

func (repo *TaskRepository) DestroyTask(id int) error {
	stmt, err := repo.Prepare("DELETE FROM task WHERE id=$1")
	if err != nil {
		log.Panic(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Panic(err)
	}
	if rowsAffected > 0 {
		return nil
	}

	return fmt.Errorf("Could not find Task with id of %d to delete", id)
}

func (repo *TaskRepository) UpdateTask(t model.Task) error {
	stmt, err := repo.Prepare("UPDATE task SET description=$1 WHERE id=$2")
	if err != nil {
		log.Panic(err)
	}
	res, err := stmt.Exec(t.Description, t.ID)
	if err != nil {
		log.Panic(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Panic(err)
	}
	if rowsAffected > 0 {
		return nil
	}

	return fmt.Errorf("Could not find Task with id of %d to update", t.ID)
}
