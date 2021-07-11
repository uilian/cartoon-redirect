package cartoon

import (
	"log"
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
	GARFIELD
	PEANUTS
	XKCD
	DEFAULT
)

func LoadCartoons() []*Cartoon {
	log.Print("Building cartoon list ...")
	cartoons := make([]*Cartoon, int(DEFAULT)+1)
	cartoons[DILBERT] = &Cartoon{DILBERT, DILBERT.toString(), time.Date(1989, 4, 16, 0, 0, 0, 0, time.UTC).Unix(), "http://dilbert.com/strip/"}
	cartoons[CALVIN] = &Cartoon{CALVIN, CALVIN.toString(), time.Date(2007, 1, 1, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/calvinandhobbes/"}
	cartoons[GARFIELD] = &Cartoon{GARFIELD, GARFIELD.toString(), time.Date(1978, 6, 19, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/garfield/"}
	cartoons[PEANUTS] = &Cartoon{PEANUTS, PEANUTS.toString(), time.Date(1950, 10, 2, 0, 0, 0, 0, time.UTC).Unix(), "https://www.gocomics.com/peanuts/"}
	cartoons[XKCD] = &Cartoon{XKCD, XKCD.toString(), 0, "https://xkcd.com/"}
	cartoons[DEFAULT] = cartoons[DILBERT]
	return cartoons
}

func (c CartoonIdx) toString() string {
	switch c {
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
	default:
		return "dilbert"
	}
}
