package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"materiel/src/redisDB"
	"materiel/src/db"
)

type Config struct {
	Mysql db.Mysql
	Redis redisDB.Redis
}

func Initial() {
	data, err := ioutil.ReadFile("config.development.toml")
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

	db.MySqlConn(&config.Mysql)
	redisDB.RedisConn(&config.Redis)
}
