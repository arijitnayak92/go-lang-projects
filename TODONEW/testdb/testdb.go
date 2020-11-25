package testdb

import (
	"database/sql"
)

// const createTable = `
// DROP TABLE IF EXISTS todo;
// Alter SEQUENCE todo_id RESTART WITH 1;
// CREATE TABLE IF NOT EXISTS todo (
//   ID int default nextval('todo_id'::regclass),
//   TITLE TEXT NOT NULL,
//   NOTE TEXT,
//   STATUS BOOLEAN
// );
// `

const createTable = `
CREATE TABLE IF NOT EXISTS todo (
  ID int default nextval('todo_id'::regclass),
  TITLE TEXT NOT NULL,
  NOTE TEXT,
  STATUS BOOLEAN
);
`

type TestDB struct {
	db *sql.DB
}

func Setup() *sql.DB {
	db, err := connectPostgresForTests()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(createTable); err != nil {
		panic(err)
	}

	return db
}

func connectPostgresForTests() (*sql.DB, error) {
	connStr := "postgres://sdxlaekjnjhlxu:b0493e3465956df4b0645747ace1a8df23377addbe929148148cebe263bd2fa5@ec2-54-157-88-70.compute-1.amazonaws.com:5432/d2nlmudsli8cmv"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
