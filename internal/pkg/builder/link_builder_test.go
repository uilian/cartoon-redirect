package builder

import (
	"regexp"
	"testing"
	"time"
)

func Test_Builder(t *testing.T) {
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
			got := BuildRedirectURL(tt.c, tt.p)
			if tt.p == "latest" && got != tt.want {
				t.Errorf("Got: '%v', wanted: '%v'", got, tt.want)
			}
			if tt.p == "random" && re.FindAllString(got, -1) == nil {
				t.Errorf("Got: '%v', wanted: '%v'", got, tt.want)
			}
		})
	}
}
