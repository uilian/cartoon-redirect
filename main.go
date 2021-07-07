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
	CALVIN
	DEFAULT
)

var (
	cartoons = map[CartoonIdx]*Cartoon{}
)

func getTargetDate(c Cartoon, period string) time.Time {
	switch {
	case period == "today":
		return time.Now()
	case period == "sunday":
		return time.Now()
	default:
		return randate(c.MinDate)
	}
}

func main() {
	loadCartoons()
	http.HandleFunc("/", cartoonHandler)
	port := os.Getenv("PORT")
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func cartoonHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("r.URL.Path: ", r.URL.Path)
	log.Print("r.URL.Query(): ", r.URL.Query())
	path := strings.Split(r.URL.Path, "/")[1]
	log.Print("path: ", path)
	period := r.URL.Query().Get("period")
	http.Redirect(w, r, buildRedirectURL(path, period), http.StatusSeeOther)
}

func loadCartoons() {
	cartoons[DILBERT] = &Cartoon{DILBERT, DILBERT.toString(), time.Date(1989, 4, 16, 0, 0, 0, 0, time.UTC).Unix(), "http://dilbert.com/strip/"}
	cartoons[CALVIN] = &Cartoon{CALVIN, CALVIN.toString(), time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/calvinandhobbes/"}
	cartoons[DEFAULT] = cartoons[DILBERT]
}

func cartoonSelector(name string) Cartoon {
	for _, v := range cartoons {
		if v.Name == name {
			return *v
		}
	}
	return *(cartoons[DEFAULT])
}

func buildRedirectURL(name string, period string) string {
	c := cartoonSelector(name)
	switch c.ID {
	case DILBERT:
		return dilbertURL(c, getTargetDate(c, period))
	case CALVIN:
		return calvinURL(c, getTargetDate(c, period))
	default:
		return dilbertURL(c, getTargetDate(c, period))
	}
}

func dilbertURL(c Cartoon, t time.Time) string {
	log.Print(c, t)
	// target := today.AddDate(-5, 0, 0)
	path := c.BaseUrl + t.Format("2006-01-02")
	return path
}

func calvinURL(c Cartoon, t time.Time) string {
	log.Print(c, t)
	path := c.BaseUrl + t.Format("2006/01/02")
	return path
}

func randate(min int64) time.Time {
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func (c CartoonIdx) toString() string {
	switch c {
	case DILBERT:
		return "dilbert"
	case CALVIN:
		return "calvin"
	default:
		return "dilbert"
	}
}
