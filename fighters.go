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

func FighterWins(id string) error {
	conn := pool.Get()
	defer conn.Close()

	exists, err := redis.Int(conn.Do("EXISTS", "fighter:"+id))
	if err != nil {
		return err
	} else if exists == 0 {
		return ErrNoFighter
	}
	_, err = conn.Do("HINCRBY", "fighter:"+id, "W", 1)
	if err != nil {
		return err
	}
	return nil
}

func FighterLose(id string) error {
	conn := pool.Get()
	defer conn.Close()

	exists, err := redis.Int(conn.Do("EXISTS", "fighter:"+id))
	if err != nil {
		return err
	} else if exists == 0 {
		return ErrNoFighter
	}
	_, err = conn.Do("HINCRBY", "fighter:"+id, "L", 1)
	if err != nil {
		return err
	}
	return nil
}
