package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	Conn()
}

func Conn() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(localhost)/materiel")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connect Mysql Successify!")
}

func Close() {
	DB.Close()
}
