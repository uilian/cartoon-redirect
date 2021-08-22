package cartoon

import (
	"log"
	"math/rand"
	"sync"
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
	UNKNOWN = CartoonIdx(iota)
	DILBERT
	CALVIN
	GARFIELD
	PEANUTS
	XKCD
	end
)

var once sync.Once
var cartoonList *[]Cartoon

func GetCartoonList() *[]Cartoon {
	once.Do(func() {
		cartoonList = loadCartoons()
	})
	return cartoonList
}

func loadCartoons() *[]Cartoon {
	log.Print("Building cartoon list ...")
	cartoons := make([]Cartoon, int(end))
	cartoons[UNKNOWN] = Cartoon{DILBERT, DILBERT.toString(), time.Date(1989, 4, 16, 0, 0, 0, 0, time.UTC).Unix(), "http://dilbert.com/strip/"}
	cartoons[DILBERT] = Cartoon{DILBERT, DILBERT.toString(), time.Date(1989, 4, 16, 0, 0, 0, 0, time.UTC).Unix(), "http://dilbert.com/strip/"}
	cartoons[CALVIN] = Cartoon{CALVIN, CALVIN.toString(), time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/calvinandhobbes/"}
	cartoons[GARFIELD] = Cartoon{GARFIELD, GARFIELD.toString(), time.Date(1978, 6, 19, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/garfield/"}
	cartoons[PEANUTS] = Cartoon{PEANUTS, PEANUTS.toString(), time.Date(1950, 10, 2, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/peanuts/"}
	cartoons[XKCD] = Cartoon{XKCD, XKCD.toString(), 0, "https://xkcd.com/"}
	return &cartoons
}

func Random() Cartoon {
	return (*GetCartoonList())[rand.Intn(int(end))]
}

func (c CartoonIdx) isValid() bool {
	return c < end
}

func (c CartoonIdx) toString() string {
	switch c {
	default:
		return "unknown"
	case DILBERT:
		return "dilbert"
	case CALVIN:
		return "calvin"
	case GARFIELD:
		return "garfield"
	case PEANUTS:
		return "peanuts"
	case XKCD:
		return "xkcd"
	}
}
