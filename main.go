package main

import (
	"database/sql"
	"github.com/amirazad1/simple-store/api"
	"github.com/amirazad1/simple-store/service"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	dbDriver = "mysql"
	dbSource = "root:secret@tcp(localhost:3306)/store?charset=utf8&parseTime=True&loc=Local"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	store := service.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(address)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
