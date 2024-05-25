package storage

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	once   sync.Once
	config = mysql.Config{
		User:   "back",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DB_NAME"),
	}
)

func NewDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}

		if err = db.Ping(); err != nil {
			log.Fatal(err)
		}
	})
}
