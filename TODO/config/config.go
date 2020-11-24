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
			host:     "ec2-54-157-88-70.compute-1.amazonaws.com",
			port:     5432,
			name:     "d2nlmudsli8cmv",
			user:     "sdxlaekjnjhlxu",
			password: "b0493e3465956df4b0645747ace1a8df23377addbe929148148cebe263bd2fa5",
			sslmode:  "disabled",
		},
	}
}

func Db() DBConfig {
	return appConfig.db
}
