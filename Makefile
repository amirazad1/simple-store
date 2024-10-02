mysql:
    docker run --name mysql-server -p 3306:3306 -e MY_SQL_ROOT_PASSWORD=secret -d mysql

createdb:
    docker exec -it mysql-server mysql -u root -e "create database store;"

dropdb:
    docker exec -it mysql-server mysql -u root -e "drop database store;"

migrateup:
    migrate -path db/migration -database "mysql://root:@tcp(localhost:3306)/store" -verbose up

migratedown:
    migrate -path db/migration -database "mysql://root:@tcp(localhost:3306)/store" -verbose down

sqlc:
    sqlc generate

test:
    go test -v -cover ./...

getall:
    go get -u -v -f all

.PHONY:mysql createdb dropdb migrateup migratedown sqlc test