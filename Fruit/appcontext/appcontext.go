package appcontext

// AppContext contains all global contexts required by application
type AppContext struct {
	PostgresURI string
}

// NewAppContext returns new instance of AppContext
func NewAppContext(postgresURI string) *AppContext {
	return &AppContext{PostgresURI: postgresURI}
}
