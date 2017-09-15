package repo

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/bgiegel/TodoREST/model"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var november2017 = time.Date(2017, time.November, 10, 23, 0, 0, 0, time.UTC)

var taskRepo *TaskRepository

func TestFindTaskToReturnTheTask(t *testing.T) {
	//given
	givenOneRowIsFoundInDB(t)
	defer taskRepo.DB.Close()
	expectedTask := model.Task{ID: 4, Creation: november2017, Description: "Description"}

	//when
	actualTask := taskRepo.FindTask(4)

	//then
	assert.Equal(t, expectedTask, actualTask)
}

func TestFindTaskToReturnEmptyTaskWhenNotFound(t *testing.T) {
	//given
	givenNoRowsAreFoundInDB(t)
	defer taskRepo.DB.Close()
	expectedTask := model.Task{}

	//when
	actualTask := taskRepo.FindTask(4)

	//then
	assert.Equal(t, expectedTask, actualTask)
}

func TestFindTasksToPanicWhenUnknonwError(t *testing.T) {
	//given
	givenUnknowErrorDuringSQLRequest(t)
	defer taskRepo.DB.Close()

	//when
	assert.Panics(t, func() { taskRepo.FindTask(1) }, "Did not panic")
}

func TestAllTasksToReturnAllStoredTasks(t *testing.T) {
	//given
	givenRowsAreFoundInDB(t)
	defer taskRepo.DB.Close()
	expectedTasks := model.Tasks{
		model.Task{ID: 1, Creation: november2017, Description: "Task1"},
		model.Task{ID: 2, Creation: november2017, Description: "Task2"},
	}

	//when
	actualTasks := taskRepo.AllTasks()

	//then
	assert.Equal(t, expectedTasks, actualTasks)
}

func TestAllTasksToReturnNoTasksWhenNoRowsFound(t *testing.T) {
	//given
	givenNoRowsAreFoundInDB(t)
	defer taskRepo.DB.Close()
	expectedTasks := model.Tasks{}

	//when
	actualTasks := taskRepo.AllTasks()

	//then
	assert.Equal(t, expectedTasks, actualTasks)
}

func TestAllTasksToPanicWhenUnknonwError(t *testing.T) {
	//given
	givenUnknowErrorDuringSQLRequest(t)
	defer taskRepo.DB.Close()

	//when
	assert.Panics(t, func() { taskRepo.AllTasks() }, "Did not panic")
}

func TestCreateTaskToReturnIdOfCreatedTask(t *testing.T) {
	//given
	givenInsertionSucceeds(t)
	defer taskRepo.DB.Close()
	taskToCreate := model.Task{Description: "Description"}

	//when
	taskID := taskRepo.CreateTask(taskToCreate)

	assert.Equal(t, 1, taskID)
}

func TestCreateTaskToPanicWhenErrorOccursDuringRequest(t *testing.T) {
	//given
	givenUnknowErrorDuringSQLRequest(t)
	defer taskRepo.DB.Close()

	//when
	assert.Panics(t, func() { taskRepo.CreateTask(model.Task{}) }, "Did not panic")
}

func TestDestroyTaskToReturnNilWhenDeleteSucceeds(t *testing.T) {
	//given
	givenRequestSucceeds(t, "DELETE")

	//when
	err := taskRepo.DestroyTask(3)

	//then
	assert.Nil(t, err)
}

func TestDestroyTaskToPanicWhenDeleteFails(t *testing.T) {
	//given
	givenRequestFails(t, "DELETE")

	//then
	assert.Panics(t, func() { taskRepo.DestroyTask(3) }, "Did not panic")
}

func TestDestroyTaskToReturnErrorWhenNoRowsAreDeleted(t *testing.T) {
	//given
	givenRequestAffectNoRows(t, "DELETE")

	//when
	actualError := taskRepo.DestroyTask(3)

	//then
	expectedErr := fmt.Errorf("Could not find Task with id of %d to delete", 3)
	assert.Equal(t, expectedErr, actualError)
}

func TestUpdateTaskToReturnNilWhenUpdateSucceeds(t *testing.T) {
	//given
	givenRequestSucceeds(t, "UPDATE")

	//when
	err := taskRepo.UpdateTask(model.Task{})

	//then
	assert.Nil(t, err)
}

func TestUpdateTaskToPanicWhenUpdateFails(t *testing.T) {
	//given
	givenRequestFails(t, "UPDATE")

	//then
	assert.Panics(t, func() { taskRepo.UpdateTask(model.Task{}) }, "Did not panic")
}

func TestUpdateTaskToReturnErrorWhenNoRowsAreUpdated(t *testing.T) {
	//given
	givenRequestAffectNoRows(t, "UPDATE")

	//when
	actualError := taskRepo.UpdateTask(model.Task{ID: 3})

	//then
	expectedErr := fmt.Errorf("Could not find Task with id of %d to update", 3)
	assert.Equal(t, expectedErr, actualError)
}

func givenOneRowIsFoundInDB(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "creation", "description"}).
		AddRow(4, november2017, "Description")

	mock := buildFakeTaskRepo(t)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
}

func givenRowsAreFoundInDB(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "creation", "description"}).
		AddRow(1, november2017, "Task1").
		AddRow(2, november2017, "Task2")

	mock := buildFakeTaskRepo(t)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
}

func givenNoRowsAreFoundInDB(t *testing.T) {
	mock := buildFakeTaskRepo(t)
	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)
}

func givenUnknowErrorDuringSQLRequest(t *testing.T) {
	mock := buildFakeTaskRepo(t)
	mock.ExpectQuery("SELECT").WillReturnError(sql.ErrTxDone)
}

func givenInsertionSucceeds(t *testing.T) {
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	mock := buildFakeTaskRepo(t)
	mock.ExpectQuery("INSERT").WillReturnRows(rows)
}

func givenRequestSucceeds(t *testing.T, execType string) {
	mock := buildFakeTaskRepo(t)
	mock.ExpectPrepare(execType).ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
}

func givenRequestFails(t *testing.T, execType string) {
	mock := buildFakeTaskRepo(t)
	mock.ExpectPrepare(execType).ExpectExec().WillReturnError(sql.ErrTxDone)
}

func givenRequestAffectNoRows(t *testing.T, execType string) {
	mock := buildFakeTaskRepo(t)
	mock.ExpectPrepare(execType).ExpectExec().WillReturnResult(sqlmock.NewResult(1, 0))
}

func buildFakeTaskRepo(t *testing.T) sqlmock.Sqlmock {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := &Repository{db}
	taskRepo = &TaskRepository{repo}
	return mock
}
