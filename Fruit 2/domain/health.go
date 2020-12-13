package domain

// CheckDatabaseHealth ...
func (d *Domain) CheckDatabaseHealth() (error, error) {
	err := d.appDB.PingPostgres()
	err2 := d.appDB.CheckMongoAlive()

	if err != nil {
		return err, nil
	}

	if err2 != nil {
		return nil, err2
	}

	return nil, nil
}
