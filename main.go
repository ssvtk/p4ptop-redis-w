package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	//ImportData("data.json")
	http.HandleFunc("/fighter", showFighter)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "hello") })
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//showFighter ...
func showFighter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, http.StatusText(405), 405)
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	fighter, err := FindFighter(id)
	if err == ErrNoFighter {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Name: %s\nW: %d L: %d\nDivision: %s", fighter.Name, fighter.Wins, fighter.Loses, fighter.Division)

}
