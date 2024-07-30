package storage

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewDB() {
	once.Do(func() {
		var (
			err    error
			config = mysql.Config{
				User:   os.Getenv("MYSQL_USER"),
				Passwd: os.Getenv("MYSQL_PASSWORD"),
				Net:    os.Getenv("MYSQL_NET"),
				Addr:   os.Getenv("MYSQL_ADDRESS"),
				DBName: os.Getenv("MYSQL_DATABASE"),
			}
		)

		db, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	})
}
