package timeparse

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:13", true},
		{"1:3:44", true},
		{"bad", false},
		{"", false},
		{"19:00", false},
		{"aa:bb:cc", false},
		{"19:00:", false},
		{":00:13", false},
		{"11::13", false},
		{"24:59:59", false},
		{"23:60:59", false},
		{"23:59:60", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
