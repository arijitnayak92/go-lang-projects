package domain

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDBConn() (con *pg.DB) {
	address := fmt.Sprintf("%s:%s", "localhost", "5432")
	options := &pg.Options{
		User:     "postgres",
		Password: "root",
		Addr:     address,
		Database: "postgres",
		PoolSize: 50,
	}
	con = pg.Connect(options)
	if con == nil {
		log.Error("cannot connect to postgres")
	}
	return
}
