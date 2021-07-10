package datastore

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// DataStore provides mongodb datastore
type DataStore struct{}

// Init initializes connection to mongodb server
func Init() *sqlx.DB {
	val, exist := os.LookupEnv("DATABASE_URL")
	if !exist {
		log.Fatalln("DATABASE_URL value not set")
	}

	datasource, err := sqlx.Connect("mysql", val)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Connection successfull")

	return datasource
}
