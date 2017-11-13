package db

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	Mysql Mysql
}

type Mysql struct {
	Host     string
	User     string
	Password string
	Port     int64
}

func init() {
	data, err := ioutil.ReadFile("config.toml")
	if err != nil {
		fmt.Println("配置文件读取失败")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(0)
	}

	var config Config
	if _, err := toml.Decode(string(data), &config); err != nil {
		fmt.Println("配置文件读取失败")
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(0)
	}

	fmt.Println(config)

	MySqlConn(&config.Mysql)
}
