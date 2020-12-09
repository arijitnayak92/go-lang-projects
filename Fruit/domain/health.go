package domain

import "log"

func (d *Domain) GetPostgresHealth() bool {
	err := d.appDB.PingPostgres()
	if err != nil {
		log.Println("Cannot connect to Postgres Database", err)
		return false
	}

	return true
}

func (d *Domain) GetMongoHealth() bool {
	err := d.appDB.CheckMongoAlive()
	if err != nil {
		log.Println("Cannot connect to Postgres Database", err)
		return false
	}

	return true
}
