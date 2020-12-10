package mock

import "github.com/arijitnayak92/taskAfford/Fruit/utils"

type Postgres struct {
}

type Mongo struct {
}

type AppDB struct {
	pg    *Postgres
	mongo *Mongo
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func NewMongo() *Mongo {
	return &Mongo{}
}

func NewDB(pg *Postgres, mongo *Mongo) *AppDB {
	return &AppDB{pg: pg, mongo: mongo}
}

func (mock *AppDB) PingPostgres() *utils.APIError {
	return nil
}

func (mock *AppDB) CheckMongoAlive() *utils.APIError {
	return nil
}
