package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/arijitnayak92/taskAfford/TODONEW/schema"
	"github.com/arijitnayak92/taskAfford/TODONEW/utils"
	_ "github.com/lib/pq"
)

//Postgres...(
type Postgres struct {
	DB *sql.DB
}

//...
func (p *Postgres) Close() {
	p.DB.Close()
}

//...
func (p *Postgres) Insert(todo *schema.Todo) (int, *utils.APIError) {
	query := "INSERT INTO todo (id, title, note, status) VALUES (nextval('todo_id'), ?, ?, ?) RETURNING id"

	rows, err := p.DB.Query(query, todo.Title, todo.Note, todo.Status)
	fmt.Println(err)
	if err != nil {
		return -1, &utils.APIError{
			Message:    "Error in DB structure !",
			StatusCode: 422,
		}
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, &utils.APIError{
				Message:    "No Relation Found !",
				StatusCode: 404,
			}
		}
	}

	return id, nil
}

//...
func (p *Postgres) GetOne(id int) (schema.Todo, *utils.APIError) {
	var t schema.Todo
	query := "SELECT id, title, note, status FROM todo WHERE id = ?"

	row := p.DB.QueryRow(query, id)

	if err := row.Scan(&t.ID, &t.Title, &t.Note, &t.Status); err != nil {
		fmt.Println(err)
		return schema.Todo{}, &utils.APIError{
			Message:    "Datas Not Found !",
			StatusCode: 404,
		}
	}

	return t, nil
}

//...
func (p *Postgres) Delete(id int) *utils.APIError {
	_, notFound := p.GetOne(id)
	fmt.Println(notFound)
	if notFound != nil {
		return &utils.APIError{
			Message:    "No Data Found Against this Id !",
			StatusCode: 404,
		}
	}
	query := "DELETE FROM todo WHERE id = ?"

	if _, err := p.DB.Exec(query, id); err != nil {
		fmt.Println(err)
		return &utils.APIError{
			Message:    "DB Execuation Failed !",
			StatusCode: 422,
		}
	}

	return nil
}

//...
func (p *Postgres) Update(id int, todo *schema.Todo) *utils.APIError {
	_, notFound := p.GetOne(id)
	if notFound != nil {
		return &utils.APIError{
			Message:    "No Data Found Against this Id !",
			StatusCode: 404,
		}
	}

	query := "UPDATE todo SET title = ?,note = ? WHERE id = ?"

	if _, err := p.DB.Exec(query, todo.Title, todo.Note, id); err != nil {
		fmt.Println(err)
		return &utils.APIError{
			Message:    "DB Execuation Failed !",
			StatusCode: 422,
		}
	}

	return nil
}

//...
func (p *Postgres) MarkAsDone(id int, status bool) *utils.APIError {
	_, notFound := p.GetOne(id)
	if notFound != nil {
		return &utils.APIError{
			Message:    "No Data Found Against this Id !",
			StatusCode: 404,
		}
	}

	query := `
        UPDATE todo SET status = $1 WHERE id = $2
    `

	if _, err := p.DB.Exec(query, status, id); err != nil {
		return &utils.APIError{
			Message:    "DB Execuation Failed !",
			StatusCode: 422,
		}
	}

	return nil
}

//...
func (p *Postgres) GetAll() ([]schema.Todo, *utils.APIError) {
	query := `
        SELECT *
        FROM todo
        ORDER BY id;
	`

	rows, err := p.DB.Query(query)

	if err != nil {
		return nil, &utils.APIError{
			Message:    "Wrong Schema Structure !",
			StatusCode: 422,
		}
	}

	var todoList []schema.Todo
	for rows.Next() {
		var t schema.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Note, &t.Status); err != nil {
			return nil, &utils.APIError{
				Message:    "Datas Not Found !",
				StatusCode: 404,
			}
		}
		todoList = append(todoList, t)
	}

	return todoList, nil
}

//...
func ConnectPostgres() (*Postgres, *utils.APIError) {
	connStr := os.Getenv("POSTGRES_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Unable to connect DB !",
			StatusCode: 400,
		}
	}

	err = db.Ping()
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Failed to ping DB !",
			StatusCode: 400,
		}
	}

	return &Postgres{db}, nil
}
