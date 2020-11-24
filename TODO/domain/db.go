package domain

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	// connString := "user=postgres dbname=postgres password=root host=localhost sslmode=disable"
	db, err = sql.Open("postgres", "postgres://sdxlaekjnjhlxu:b0493e3465956df4b0645747ace1a8df23377addbe929148148cebe263bd2fa5@ec2-54-157-88-70.compute-1.amazonaws.com:5432/d2nlmudsli8cmv")
	fmt.Println("Going to connect !")
	fmt.Println(db)
	if err != nil {
		fmt.Println("Database connection failed !")
		log.Fatalf("failed to load the database: %s", err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Database ping connection failed !")
		log.Fatalf("ping to database failed :%s", err)
	}
	return db
}
