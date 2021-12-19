package cartoon

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

func TestGetCartoonList(t *testing.T) {
	cl := GetCartoonList()
	l := len(*cl)
	if l > int(end) {
		t.Errorf("got len(cartoons) = %d; wanted %d", l, int(end))
	}
	if (*cl)[DILBERT] != (*cl)[UNKNOWN] {
		t.Errorf("got wrong default value %v; wanted %v", (*cl)[UNKNOWN], (*cl)[DILBERT])
	}

	for _, c := range *cl {
		if got := c.id; !got.isValid() {
			t.Errorf("cartoon id is not valid = %v, want %v", got, "0.."+fmt.Sprint(end-1))
		}
	}
}

func TestURL(t *testing.T) {
	var re = regexp.MustCompile(`(?m)^(ftp|http|https):\/\/[^ "]+$`)
	tests := []struct {
		name string
		c, p string
		want string
	}{
		// {"empty", "", "", ""},
		{"dibert empty period", "dilbert", "", "http://dilbert.com/strip/" + time.Now().Format("2006-01-02")},
		{"random period", "", "random", "Valid URL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateURL(tt.c, tt.p)
			if tt.p == "latest" && got != tt.want {
				t.Errorf("Got: '%v', wanted: '%v'", got, tt.want)
			}
			if tt.p == "random" && re.FindAllString(got, -1) == nil {
				t.Errorf("Got: '%v', wanted: '%v'", got, tt.want)
			}
		})
	}
}
