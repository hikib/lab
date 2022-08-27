package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("Hiko")
	want := "Hello, Hiko"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
