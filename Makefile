dbcontainer: 
	docker run --name moggdb -e MYSQL_ROOT_PASSWORD=secret -p 3307:3306 -d mysql:latest

createdb:
	docker exec -it moggdb mysql -u root -p"secret" -e "CREATE DATABASE ismogged;" 

dropdb:
	docker exec -it moggdb mysql -u root -p"secret" -e "DROP DATABASE ismogged;" 

mysqlshell:
	docker exec -it moggdb bash

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(127.0.0.1:3307)/ismogged?parseTime=true" --verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(127.0.0.1:3307)/ismogged" --verbose down 

sqlc:
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate

ctest:
	go clean -testcache && go test -v -cover ./...

coveragereport:
	go test -coverprofile=coverage/coverage.out ./... && go tool cover -html=coverage/coverage.out -o=coverage/coverage.html

.PHONY: dbcontainer createdb dropdb mysqlshell migrateup migratedown sqlc ctest
