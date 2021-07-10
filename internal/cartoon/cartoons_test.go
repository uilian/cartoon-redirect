package cartoon

import "testing"

func Test_LoadCartoons(t *testing.T) {
	c := LoadCartoons()
	l := len(c)
	if l != int(DEFAULT+1) {
		t.Errorf("len(cartoons) = %d; want %d", l, int(DEFAULT+1))
	}
	if c[DILBERT] != c[DEFAULT] {
		t.Errorf("wrong default value, should be = %v; got %v", *c[DILBERT], *c[DEFAULT])
	}
}
