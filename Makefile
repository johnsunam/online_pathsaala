create-new-migrate:
	migrate create -ext sql -dir dbs/migrate -seq $(action)

migrate-up:
	migrate -path dbs/migrate/ -database "postgresql://postgres:12345@localhost:5432/online-pathsaala?sslmode=disable" -verbose up

run-app: 
	air
