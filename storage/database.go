package storage

import (
	"database/sql"
	"fmt"
	"log"
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
		DBName: "gamer_shop",
	}
)

func NewDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			log.Fatalf(err.Error())
		}

		if err = db.Ping(); err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println("Connect to MySQL")
	})
}
