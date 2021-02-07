package appcontext

// AppContext ...
type AppContext struct {
	PostgresURI string
	MongoURI    string
}

// NewAppContext ...
func NewAppContext(postgresURI string, mongoURI string) *AppContext {
	return &AppContext{PostgresURI: postgresURI, MongoURI: mongoURI}
}
