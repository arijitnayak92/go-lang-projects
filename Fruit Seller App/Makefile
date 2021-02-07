init-db:
	psql -c "create database fruitseller" -U $(user)

migration-init:
	migrate create -ext sql -dir db/migrations -seq fs-postgres

migration-mongo-init:
	migrate create -ext sql -dir db/migration -seq fs-mongo
