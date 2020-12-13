package appcontext

// AppContext ...
type AppContext struct {
	PostgresURI string
	MongoURI    string
}

// NewAppContext returns new instance of AppContext
func NewAppContext(postgresURI string, mongoURI string) *AppContext {
	return &AppContext{PostgresURI: postgresURI, MongoURI: mongoURI}
}
