package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "mysql"
	dbSource = "root:secret@tcp(localhost:3306)/store?charset=utf8&parseTime=True&loc=Local"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
