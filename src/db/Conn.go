package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func MySqlConn(config *Mysql) {
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/materiel", config.User, config.Password, config.Host))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connect Mysql Successify!")
}

func Close() {
	DB.Close()
}
