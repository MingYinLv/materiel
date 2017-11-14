package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Mysql struct {
	Host     string
	User     string
	Password string
	Port     int64
	DB       string
}


func MySqlConn(config *Mysql) {
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.Host, config.DB))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connect Mysql Successify!")
}

func Close() {
	DB.Close()
}
