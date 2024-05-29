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
				User:   os.Getenv("DB_USER"),
				Passwd: os.Getenv("DB_PASS"),
				Net:    os.Getenv("DB_NET"),
				Addr:   os.Getenv("DB_ADDR"),
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
