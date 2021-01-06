package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ImportData(filename string) {
	conn := pool.Get()
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var fighters Fighters
	value, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(value, &fighters)
	if err != nil {
		log.Fatal(err)
	}
	for i, fighter := range fighters.Fighters {
		fmt.Println(fighter.Name, fighter.Wins, fighter.Loses, fighter.Division)
		_, err := conn.Do("HMSET",
			fmt.Sprintf("fighter:%d", i+1),
			"name", fighter.Name,
			"W", fighter.Wins,
			"L", fighter.Loses,
			"division", fighter.Division)
		if err != nil {
			log.Fatal(err)
		}
	}

}
