package service

import (
	"database/sql"
	db "github.com/amirazad1/simple-store/db/sqlc"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "mysql"
	dbSource = "root:@tcp(localhost:3306)/store?charset=utf8&parseTime=True&loc=Local"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	testQueries = db.New(testDB)
	os.Exit(m.Run())
}
