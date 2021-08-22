package cartoon

import (
	"fmt"
	"testing"
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
		if got := c.ID; !got.isValid() {
			t.Errorf("cartoon id is not valid = %v, want %v", got, "0.."+fmt.Sprint(end-1))
		}
	}
}
