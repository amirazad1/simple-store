package db

import (
	"database/sql"
	"github.com/amirazad1/simple-store/util"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load config file", err)
	}
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
