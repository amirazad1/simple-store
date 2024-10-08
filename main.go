package main

import (
	"database/sql"
	"github.com/amirazad1/simple-store/api"
	"github.com/amirazad1/simple-store/service"
	"github.com/amirazad1/simple-store/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config file", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	store := service.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can not create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}
}
