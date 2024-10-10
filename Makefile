createnetwork:
	docker network create store-network
mysql:
	docker run --name mysql-server --network store-network -p 3306:3306 -e MY_SQL_ROOT_PASSWORD=secret -d mysql

createdb:
	docker exec -it mysql-server mysql -u root -e "create database store;"

dropdb:
	docker exec -it mysql-server mysql -u root -e "drop database store;"

migrateup:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/store" -verbose up

migrateup1:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/store" -verbose up 1

migratedown:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/store" -verbose down

migratedown1:
	migrate -path db/migration -database "mysql://root:secret@tcp(localhost:3306)/store" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

getall:
	go get -u -v -f all

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/amirazad1/simple-store/service Store

app_image:
	docker build -t simplestore:latest .

app:
	docker run --name:simple-store --network store-network -p 8080:8080 -e GIN_MODE=release DB_SOURCE = "root:secret@tcp(mysql-server:3306)/store?charset=utf8&parseTime=True&loc=Local" simplestore:latest

.PHONY:createnetwork mysql createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock app_image app