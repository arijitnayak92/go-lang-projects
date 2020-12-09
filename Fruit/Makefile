init-db:
	psql -c "create database affordmed" -U $(user)

# You will have to install migrate package: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
migration-init:
	migrate create -ext sql -dir db/migrations -seq $(name)

# DatabaseURL=postgresql://[user]:[password]@[ip]:[port]
migration-up:
	migrate -path db/migrations -database "$(databaseURL)/affordmed?sslmode=disable" -verbose up

# DatabaseURL=postgresql://[user]:[password]@[ip]:[port]
migration-down:
	migrate -path db/migrations -database "$(databaseURL)/affordmed?sslmode=disable" -verbose down

# version=int(version to fix - 1)
migration-fix:
	migrate -path db/migrations -database "$(databaseURL)/affordmed?sslmode=disable" force $(version)