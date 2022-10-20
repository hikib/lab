package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Hiko", "")
		want := "Hello, Hiko"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty input", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloLanguages(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Hiko", "Spanish")
		want := "Hola, Hiko"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Hiko", "German")
		want := "Hallo, Hiko"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
