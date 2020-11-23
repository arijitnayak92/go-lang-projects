package config

var appConfig Config

type Config struct {
	port       int
	appName    string
	appVersion string
	db         DBConfig
}

func Load() {
	appConfig = Config{
		db: DBConfig{
			host:     "localhost",
			port:     5432,
			name:     "todo",
			user:     "postgres",
			password: "root",
		},
	}
}

func Db() DBConfig {
	return appConfig.db
}
