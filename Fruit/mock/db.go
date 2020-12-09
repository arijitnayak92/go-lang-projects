package mock

import "database/sql"

type Postgres struct {
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Exec(_ string, _ ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (p *Postgres) QueryRow(_ string, _ ...interface{}) *sql.Row {
	return nil
}

func (p *Postgres) Ping() error {
	return nil
}
