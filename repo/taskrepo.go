package repo

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bgiegel/TodoREST/model"
)

var currentID int

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("postgres", "postgres://todo_user:password@localhost/tododb")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if db != nil {
		log.Println("Connection to tododb -> OK")
	}
}

func RepoFindTask(id int) model.Task {
	task := model.Task{}
	row := db.QueryRow("SELECT * FROM task WHERE id=$1", id)
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

func RepoAllTasks() (model.Tasks, error) {
	log.Println("Returning tasks....")

	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]model.Task, 0)
	for rows.Next() {
		task := model.Task{}
		err := rows.Scan(&task.ID, &task.Creation, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	response := tasks

	return response, nil
}

func RepoCreateTask(t model.Task) (taskID int) {
	err := db.QueryRow("INSERT INTO task(description) VALUES($1) RETURNING id", t.Description).Scan(&taskID)
	if err != nil {
		log.Panic(err)
	}
	return
}

func RepoDestroyTask(id int) error {
	stmt, err := db.Prepare("DELETE FROM task WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rowsAffected > 0 {
		return nil
	}

	return fmt.Errorf("Could not find Task with id of %d to delete", id)
}

func RepoUpdateTask(t model.Task) error {
	stmt, err := db.Prepare("UPDATE task SET description=$1 WHERE id=$2")
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

	return fmt.Errorf("Could not find Task with id of %d to delete", t.ID)
}
