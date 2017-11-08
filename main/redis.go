package main

import (
	"fmt"
	"materiel/src/redis"
)

func main() {
	fmt.Println(redis.Get("name"))
}
