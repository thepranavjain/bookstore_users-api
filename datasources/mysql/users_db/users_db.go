package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	// username, password, host, schema, charset from config.go in this package
	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		username, password, host, schema, charset,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to mysql users database")
}
