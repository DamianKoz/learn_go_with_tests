package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}
	t.Run("Name is given", func(t *testing.T) {
		got := Hello("Jack", "")
		want := "Hello, Jack"

		assertCorrectMessage(t, got, want)
	})
	t.Run("No Name given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jack", "Spanish")
		want := "Hola, Jack"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jack", "French")
		want := "Bonjour, Jack"
		assertCorrectMessage(t, got, want)
	})
}
