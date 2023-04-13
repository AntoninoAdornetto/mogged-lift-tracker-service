dbcontainer: 
	docker run --name moggdb -e MYSQL_ROOT_PASSWORD=secret -p 3307:3306 -d mysql:latest

createdb:
	docker exec -it moggdb mysql -u root -p"secret" -e "CREATE DATABASE ismogged;" 

dropdb:
	docker exec -it moggdb mysql -u root -p"secret" -e "DROP DATABASE ismogged;" 

mysqlshell:
	docker exec -it moggdb bash

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3307)/ismogged" --verbose up

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3307)/ismogged" --verbose down 

.PHONY: dbcontainer createdb dropdb mysqlshell migrateup
