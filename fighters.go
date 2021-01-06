package main

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

//Pool of redis connections
var pool *redis.Pool

var ErrNoFighter = errors.New("no album found")

//Fighter ...
type Fighter struct {
	Name string `redis:"name"`
	Wins int `redis:"W"`
	Loses int `redis:"L"`
	Devision string `redis:"devision"`
}


