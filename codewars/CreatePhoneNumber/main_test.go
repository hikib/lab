package main

import (
	"fmt"
	"testing"
)

func TestCreatePhoneNumber(t *testing.T) {
	var tests = []struct {
		n    string
		s    [10]uint
		want string
	}{
		{"example", [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "(123) 456-7890"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.n)
		t.Run(testname, func(t *testing.T) {
			ans := CreatePhoneNumber(tt.s)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
