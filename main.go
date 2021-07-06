package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

type (
	Cartoon int
	cartoon struct {
		ID      Cartoon
		Name    string
		MinDate int64
		BaseUrl string
	}
)

const (
	Dilbert Cartoon = iota
)

var (
	cartoons = map[Cartoon]*cartoon{}
)

func loadCartoons() {
	c := &cartoon{
		ID:      Dilbert,
		Name:    "dilbert",
		MinDate: time.Date(1989, 4, 16, 0, 0, 0, 0, time.UTC).Unix(),
		BaseUrl: "http://dilbert.com/strip/",
	}
	cartoons[Dilbert] = c
}

func main() {
	http.HandleFunc("/cartoon", cartoonHandler)
	http.HandleFunc("/dilbert", cartoonHandler)
	loadCartoons()
	http.ListenAndServe(":3000", nil)
}

func cartoonHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	http.Redirect(w, r, dilbertURL(1), http.StatusSeeOther)
}

func dilbertURL(offset int) string {
	c := cartoons[Dilbert]
	log.Print(c)
	// today := time.Now()
	// target := today.AddDate(-5, 0, 0)
	target := randate(c.MinDate)
	path := c.BaseUrl + target.Format("2006-01-02")
	return path
}

func randate(min int64) time.Time {
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
