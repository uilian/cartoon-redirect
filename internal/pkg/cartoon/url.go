package cartoon

import (
	"log"
	"math/rand"
	"time"
)

func GenerateURL(name string, period string) string {
	c := cartoonSelector(name)
	log.Print("Selected: ", c.id, " - ", c.name)
	switch c.id {
	default:
		return dilbertURL(c, getTargetDate(c, period))
	case DILBERT:
		return dilbertURL(c, getTargetDate(c, period))
	case CALVIN, GARFIELD, PEANUTS:
		return gocomicsURL(c, getTargetDate(c, period))
	case XKCD:
		return xkcdURL(c, period)

	}
}

func cartoonSelector(name string) Cartoon {
	if len(name) > 0 {
		// tries to find the cartoon with the same name
		for _, v := range *GetCartoonList() {
			if v.name == name {
				return v
			}
		}
	}
	// pick one of the available cartoons
	return Random()
}

func dilbertURL(c Cartoon, t time.Time) string {
	path := c.baseUrl + t.Format("2006-01-02")
	return path
}

func gocomicsURL(c Cartoon, t time.Time) string {
	path := c.baseUrl + t.Format("2006/01/02")
	return path
}

func xkcdURL(c Cartoon, p string) string {
	switch {
	case p == "latest":
		return c.baseUrl
	case p == "random":
		return "https://c.xkcd.com/comic/random"
	default:
		return c.baseUrl
	}
}

func getTargetDate(c Cartoon, period string) time.Time {
	switch {
	case period == "latest":
		return time.Now()
	case period == "random":
		return randate(c.minDate)
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
