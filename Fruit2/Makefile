#-------------------------------------------------------------------
#For windows
#Install make from: http://gnuwin32.sourceforge.net/packages/make.htm
#Add the make.exe in bin directory to your PATH variable.
#-------------------------------------------------------------------

db-postgres-init:
	psql -c "CREATE DATABASE FruitSeller" -U $(user)

migration-postgres-init:
	migrate create -ext sql -dir db/migration -seq fs-postgres

migration-mongo-init:
	migrate create -ext sql -dir db/migration -seq fs-mongo

migrateup-postgres:
    migrate -path db/migration -database "postgresql://$(user):$(password)@$(host):$(port)/FruitSeller?sslmode=disable" -verbose up

migratedown-postgre:
    migrate -path db/migration -database "postgresql://$(user):$(password)@$(host):$(port)/FruitSeller?sslmode=disable" -verbose down


.PHONY: db-postgres-init migration-postgres-init migration-mongo-init migrateup-postgres migratedown-postgres: