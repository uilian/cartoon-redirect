package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type (
	CartoonIdx int
	Cartoon    struct {
		ID      CartoonIdx
		Name    string
		MinDate int64
		BaseUrl string
	}
)

const (
	DILBERT = CartoonIdx(iota)
	DEFAULT
)

var (
	cartoons = map[CartoonIdx]*Cartoon{}
)

func main() {
	loadCartoons()
	http.HandleFunc("/", cartoonHandler)
	http.HandleFunc("/cartoon", cartoonHandler)
	http.HandleFunc("/dilbert", cartoonHandler)
	port := os.Getenv("PORT")
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func cartoonHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	http.Redirect(w, r, buildURL(r.URL.Path), http.StatusSeeOther)
}

func loadCartoons() {
	c := &Cartoon{DILBERT, "dilbert", time.Date(1989, 4, 16, 0, 0, 0, 0, time.UTC).Unix(), "http://dilbert.com/strip/"}
	cartoons[DILBERT] = c
	cartoons[DEFAULT] = c
}

func cartoonSelector(name string) Cartoon {
	for _, v := range cartoons {
		if v.Name == name {
			return *v
		}
	}
	return *(cartoons[DEFAULT])
}

func buildURL(path string) string {
	p := strings.ReplaceAll(path, "/", "")
	c := cartoonSelector(p)
	switch c.ID {
	case DILBERT:
		return dilbertURL(c)
	default:
		return dilbertURL(c)
	}
}

func dilbertURL(c Cartoon) string {
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
