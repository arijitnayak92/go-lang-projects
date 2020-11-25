package db

import (
	"database/sql"
	"fmt"

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
	query := `
        INSERT INTO todo (id, title, note, status)
        VALUES (nextval('todo_id'), $1, $2, $3)
        RETURNING id;
    `

	rows, err := p.DB.Query(query, todo.Title, todo.Note, todo.Status)
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
func (p *Postgres) Delete(id int) *utils.APIError {
	query := `
        DELETE FROM todo
        WHERE id = $1;
    `

	if _, err := p.DB.Exec(query, id); err != nil {
		return &utils.APIError{
			Message:    "DB Execuation Failed !",
			StatusCode: 422,
		}
	}

	return nil
}

//...
func (p *Postgres) Update(id int, todo *schema.Todo) *utils.APIError {

	query := `
       UPDATE todo SET title = $1,note = $2 WHERE id = $3
		`

	if _, err := p.DB.Exec(query, todo.Title, todo.Note, id); err != nil {
		return &utils.APIError{
			Message:    "DB Execuation Failed !",
			StatusCode: 422,
		}
	}

	return nil
}

//...
func (p *Postgres) MarkAsDone(id int, status bool) *utils.APIError {
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
func (p *Postgres) GetOne(id int) (schema.Todo, *utils.APIError) {
	var t schema.Todo
	query := `
        SELECT *
        FROM todo
        WHERE id = $1;
    `

	row := p.DB.QueryRow(query, id)
	if err := row.Scan(&t.ID, &t.Title, &t.Note, &t.Status); err != nil {
		fmt.Println("Error In Scanning !")
		return schema.Todo{}, &utils.APIError{
			Message:    "Datas Not Found !",
			StatusCode: 404,
		}
	}

	return t, nil
}

//...
func ConnectPostgres() (*Postgres, *utils.APIError) {
	connStr := "postgres://sdxlaekjnjhlxu:b0493e3465956df4b0645747ace1a8df23377addbe929148148cebe263bd2fa5@ec2-54-157-88-70.compute-1.amazonaws.com:5432/d2nlmudsli8cmv"
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
