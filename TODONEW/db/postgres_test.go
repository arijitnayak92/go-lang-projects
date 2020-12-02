package db

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	_ "github.com/lib/pq"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var todo = &schema.Todo{
	ID:     1,
	Title:  "title1",
	Note:   "note1",
	Status: false,
}

//...
func TestPostgres_GetAll(t *testing.T) {
	mockDB, mock := NewMock()
	db := &Postgres{mockDB}
	defer func() {
		mockDB.Close()
	}()

	rows := sqlmock.NewRows([]string{"id", "title", "note", "status"}).
		AddRow(1, "enTitle", "/en-link", false).
		AddRow(2, "enTitle2", "/en-link2", false)
	mock.ExpectQuery("^SELECT (.+) FROM todo*").WillReturnRows(rows)

	got, _ := db.GetAll()

	var menuLinks []schema.Todo
	menuLink1 := schema.Todo{
		ID:     1,
		Title:  "enTitle",
		Status: false,
		Note:   "/en-link",
	}
	menuLinks = append(menuLinks, menuLink1)

	menuLink2 := schema.Todo{
		ID:     2,
		Title:  "enTitle2",
		Status: false,
		Note:   "/en-link2",
	}

	menuLinks = append(menuLinks, menuLink2)
	fmt.Println(got)
	// if equal(got, menuLinks) {
	// 	t.Fatalf("Want: %v, Got: %v", menuLinks, got)
	// }
}

func TestPostgres_GetOne(t *testing.T) {
	mockDB, mock := NewMock()
	db := &Postgres{mockDB}
	defer func() {
		mockDB.Close()
	}()

	querys := "SELECT id, title, note, status FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"id", "title", "note", "status"}).AddRow(todo.ID, todo.Title, todo.Note, todo.Status)

	mock.ExpectQuery(querys).WithArgs(todo.ID).WillReturnRows(rows)

	got, err := db.GetOne(todo.ID)

	if err != nil {
		t.Fatal(err)
	}

	want := &schema.Todo{
		ID:     1,
		Title:  "title1",
		Note:   "note1",
		Status: false,
	}
	if equal(got, want) {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}

//...
func TestPostgres_Insert(t *testing.T) {
	mockDB, mock := NewMock()
	db := &Postgres{mockDB}
	defer mockDB.Close()

	mock.ExpectExec("INSERT .+").WithArgs(todo.Title, todo.Note, todo.Status).WillReturnResult(sqlmock.NewResult(1, 1))

	got, err := db.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(got)

}

//...
func TestPostgres_Update(t *testing.T) {
	mockDB, mock := NewMock()
	db := &Postgres{mockDB}
	defer func() {
		mockDB.Close()
	}()

	//query := "UPDATE todo SET title = \\?, note = \\? WHERE id = \\?"
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE todo SET (.+) WHERE (.+)").WithArgs(todo.Title, todo.Note, todo.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// prep := mock.ExpectPrepare(query)
	// prep.ExpectExec().WithArgs(todo.Title, todo.Note, todo.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	updateTodo := &schema.Todo{
		Title: "Updated",
		Note:  "UP",
	}
	err := db.Update(todo.ID, updateTodo)
	if err != nil {
		t.Fatal(err)
	}
}

// //...
// func TestPostgres_Delete(t *testing.T) {
// 	mockDB, mock := NewMock()
// 	db := &Postgres{mockDB}
// 	defer func() {
// 		db.Close()
// 	}()
// 	query := "DELETE FROM todo WHERE id = \\?"

// 	mock.ExpectQuery(query).WithArgs(todo.ID)

// 	err := db.Delete(todo.ID)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// //...
// func TestPostgres_MarkAsDone(t *testing.T) {
// 	mockDB, mock := NewMock()
// 	db := &Postgres{mockDB}
// 	defer func() {
// 		db.Close()
// 	}()

// 	query := "UPDATE todo SET title = \\?, note = \\? WHERE id = \\?"

// 	prep := mock.ExpectPrepare(query)
// 	prep.ExpectExec().WithArgs(todo.Title, todo.Note, todo.ID).WillReturnResult(sqlmock.NewResult(0, 1))

// 	err := db.Update(todo.ID, todo)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

//...

func equal(got interface{}, want interface{}) bool {
	return reflect.DeepEqual(got, want)
}
