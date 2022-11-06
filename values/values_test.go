package values

import "testing"

func TestFlipValue(t *testing.T) {
	got := GetOpposite(true)
	if got {
		t.Error("Should have flipped true to false")
	}

	got = GetOpposite(false)
	if !got {
		t.Error("Should have filpped false to true")
	}
}
