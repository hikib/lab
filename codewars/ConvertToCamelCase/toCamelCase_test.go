package main

import (
	"fmt"
	"testing"
)

func TestToCamelCase(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"The_Stealth_Warrior", "TheStealthWarrior"},
		{"the-Stealth-Warrior", "theStealthWarrior"},
		{"the-stealth-warrior", "theStealthWarrior"},
		{"The-stealth-warrior", "TheStealthWarrior"},
		{"The_stealth-warrior", "TheStealthWarrior"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.s)
		t.Run(testname, func(t *testing.T) {
			ans := ToCamelCase(tt.s)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
