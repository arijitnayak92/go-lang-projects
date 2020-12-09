package domain

import "log"

func (d *Domain) GetPostgresHealth() bool {
	err := d.pg.Ping()
	if err != nil {
		log.Println("Cannot connect to Postgres Database", err)
		return false
	}

	return true
}
