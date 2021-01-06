package main

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

//Pool of redis connections
var pool *redis.Pool

var ErrNoFighter = errors.New("no fighter found")

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

//FindFighter ...
func FindFighter(id string) (*Fighter, error) {
	conn := pool.Get()
	defer conn.Close()
	val, err := redis.Values(conn.Do("HGETALL", "fighter:"+id))
	if err != nil {
		return nil, err
	} else if len(val) == 0 {
		return nil, ErrNoFighter
	}
	var fighter Fighter
	err = redis.ScanStruct(val, &fighter)
	if err != nil {
		return nil, err
	}
	return &fighter, nil
}
