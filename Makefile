network:
	docker network create mog-network

mysql: 
	docker run --name mysql --network mog-network -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest

createdb:
	docker exec -it mysql mysql -u root -p"secret" -e "CREATE DATABASE ismogged;" 

dropdb:
	docker exec -it mysql mysql -u root -p"secret" -e "DROP DATABASE ismogged;" 

mysqlshell:
	docker exec -it mysql bash

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/ismogged?parseTime=true" --verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/ismogged" --verbose down 

sqlc:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc Store

ctest:
	go clean -testcache && go test -v -cover ./...

coveragereport:
	go test -coverprofile=coverage/coverage.out ./... && go tool cover -html=coverage/coverage.out -o=coverage/coverage.html

start-server:
	go run main.go

.PHONY: network mysql createdb dropdb mysqlshell migrateup migratedown sqlc ctest start-server
