package functioanl

import (
	// ...
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/arijitnayak92/taskAfford/TODONEW/db"
	"github.com/arijitnayak92/taskAfford/TODONEW/handler"
	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/arijitnayak92/taskAfford/TODONEW/testdb"
	"github.com/arijitnayak92/taskAfford/TODONEW/utils"
	"github.com/gorilla/mux"
)

//...
func setupServer(postgres *db.Postgres) *mux.Router {
	return handler.SetUpRouting(postgres)
}

// ...
func TestGetAllTodo(t *testing.T) {
	postgres := &db.Postgres{testdb.Setup()}
	testServer := setupServer(postgres)

	todo := &schema.Todo{
		Title:  "My Task1",
		Status: false,
	}
	var err *utils.APIError

	_, err = postgres.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}
	var errs error
	req, errs := http.NewRequest(http.MethodGet, "http://localhost:8080/todos", nil)
	if errs != nil {
		t.Fatal(errs)
	}

	rec := httptest.NewRecorder()
	testServer.ServeHTTP(rec, req)

	got := strings.TrimSpace(rec.Body.String())

	want := `[{"id":1,"title":"My Task1","note":"","status":false}]`

	if got != want {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}

func TestSaveTodo(t *testing.T) {
	postgres := &db.Postgres{testdb.Setup()}
	testServer := setupServer(postgres)

	body := []byte(`{"id":1,"title":"My Task1","note":"","status":false}`)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/todo", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	testServer.ServeHTTP(rec, req)

	got := strings.TrimSpace(rec.Body.String())
	want := "1"

	if got != want {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
	var errs *utils.APIError
	gotTodo, errs := postgres.GetAll()
	if errs != nil {
		t.Fatal(errs)
	}

	wantTodo := []schema.Todo{
		{
			Title:  "My Task1",
			Status: false,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Want: %v, Got: %v\n", wantTodo, gotTodo)
	}
}

func TestDeleteTodo(t *testing.T) {
	postgres := &db.Postgres{testdb.Setup()}
	testServer := setupServer(postgres)

	todo := &schema.Todo{
		Title:  "My Task1",
		Status: false,
	}

	id, err := postgres.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	body := []byte(fmt.Sprintf(`{"id":%d}`, id))
	var errs error
	req, errs := http.NewRequest(http.MethodDelete, `http://localhost:8080/todos/${id}`, bytes.NewReader(body))
	if errs != nil {
		t.Fatal(errs)
	}

	rec := httptest.NewRecorder()
	testServer.ServeHTTP(rec, req)

	got := rec.Body.String()

	want := "Successfully Done !"

	if got != want {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}

	gotTodo, err := postgres.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(gotTodo) > 0 {
		t.Fatalf("Should return the empty slice, Got: %v\n", gotTodo)
	}
}

func TestUpdateTodo(t *testing.T) {
	postgres := &db.Postgres{testdb.Setup()}
	testServer := setupServer(postgres)

	body := []byte(`{"id":1,"title":"My Task1[Updated]","note":"New Note","status":false}`)

	req, err := http.NewRequest(http.MethodPut, `http://localhost:8080/todos/${body.id}`, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	testServer.ServeHTTP(rec, req)

	got := strings.TrimSpace(rec.Body.String())
	want := "Successfully Updated !"

	if got != want {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
	var errs *utils.APIError
	gotTodo, errs := postgres.GetAll()
	if errs != nil {
		t.Fatal(errs)
	}

	wantTodo := []schema.Todo{
		{
			Title:  "My Task1",
			Status: false,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Want: %v, Got: %v\n", wantTodo, gotTodo)
	}
}
