package main

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

//Pool of redis connections
var pool *redis.Pool

var ErrNoFighter = errors.New("no album found")

//Fighters
type Fighters struct {
	Fighters []Fighter
}

//Fighter ...
type Fighter struct {
	Name     string `redis:"name" json:"name"`
	Wins     int    `redis:"W" json:"W"`
	Loses    int    `redis:"L" json:"L"`
	Division string `redis:"division" json:"division"`
}
