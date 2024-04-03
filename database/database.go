package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewDB() {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "back:password@tcp(localhost:3306)/gamer_shop?parseTime=true")
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("connect to MySQL")
	})
}
