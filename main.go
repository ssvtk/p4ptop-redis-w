package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	ImportData(conn, "data.json")

}
