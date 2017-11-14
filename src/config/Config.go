package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"materiel/src/db"
	"materiel/src/redisDB"
	"os"
)

type Config struct {
	Mysql      db.Mysql
	Redis      redisDB.Redis
	ListenPort int64
}

func Initial(listenPort *int64) {
	args := os.Args[1:]
	fileName := "config.development.toml"
	if len(args) > 0 {
		fileName = args[0]
	}
	fmt.Printf("读取配置文件：%s\n", fileName)
	data, err := ioutil.ReadFile(fileName)
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

	if config.ListenPort == 0 {
		*listenPort = 3333
	} else {
		*listenPort = config.ListenPort
	}

	fmt.Println(config)

	db.MySqlConn(&config.Mysql)
	redisDB.RedisConn(&config.Redis)
}
