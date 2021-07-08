package api

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	cartoon "github.com/uilian/cartoon-redirect/internal/cartoon"
)

var cartoonList []cartoon.Cartoon = cartoon.LoadCartoons()

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	period := r.URL.Query().Get("q")
	url := buildRedirectURL(path, period)
	log.Print("Redirecting to: ", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func buildRedirectURL(name string, period string) string {
	c := cartoonSelector(name)
	log.Print("Selected: ", c.ID, " - ", c.Name)
	switch c.ID {
	case cartoon.DILBERT:
		return dilbertURL(c, getTargetDate(c, period))
	case cartoon.CALVIN, cartoon.GARFIELD, cartoon.PEANUTS:
		return gocomicsURL(c, getTargetDate(c, period))
	case cartoon.XKCD:
		return xkcdURL(c, period)
	default:
		return dilbertURL(c, getTargetDate(c, period))
	}
}

func cartoonSelector(name string) cartoon.Cartoon {
	if len(name) > 0 {
		// tries to find the cartoon with the same name
		for _, v := range cartoonList {
			if v.Name == name {
				return v
			}
		}
	}
	// pick one of the available cartoons
	r := cartoon.CartoonIdx(rand.Intn(int(cartoon.DEFAULT)))
	return cartoonList[r]
}

func dilbertURL(c cartoon.Cartoon, t time.Time) string {
	path := c.BaseUrl + t.Format("2006-01-02")
	return path
}

func gocomicsURL(c cartoon.Cartoon, t time.Time) string {
	path := c.BaseUrl + t.Format("2006/01/02")
	return path
}

func xkcdURL(c cartoon.Cartoon, p string) string {
	switch {
	case p == "latest":
		return c.BaseUrl
	case p == "random":
		return "https://c.xkcd.com/comic/random"
	default:
		return c.BaseUrl
	}
}

func getTargetDate(c cartoon.Cartoon, period string) time.Time {
	switch {
	case period == "latest":
		return time.Now()
	case period == "random":
		return randate(c.MinDate)
	default:
		return time.Now()
	}
}

func randate(min int64) time.Time {
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
