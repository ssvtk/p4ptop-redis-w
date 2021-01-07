package main

import (
	"github.com/gomodule/redigo/redis"
	"html/template"
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
	http.HandleFunc("/win", addWin)
	http.HandleFunc("/lose", addLose)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//showFighter ...
func showFighter(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.gohtml")

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

	err = tmpl.Execute(w, fighter)

}

//addWin ...
func addWin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
	}
	id := r.PostFormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err := FighterWins(id)
	if err == ErrNoFighter {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return

	}

	http.Redirect(w, r, "/fighter?id="+id, 303)

}

//addLose ...
func addLose(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(405), 405)
	}
	id := r.PostFormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err := FighterLose(id)
	if err == ErrNoFighter {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return

	}

	http.Redirect(w, r, "/fighter?id="+id, 303)

}
