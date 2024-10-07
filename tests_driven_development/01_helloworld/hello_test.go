package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("chris", "")
		want := "hello, chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("default to 'hello, world' when provided an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "hello, world"

		assertCorrectMessage(t, got, want)
	})
	t.Run("language passed in - Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("language passed in - French", func(t *testing.T) {
		got := Hello("kunal", "french")
		want := "Bonjour, kunal"
		assertCorrectMessage(t, got, want)
	})
	t.Run("language passed in - Hindi", func(t *testing.T) {
		got := Hello("kunal", "hindi")
		want := "Namaste, kunal"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // to tell that this method is a "helper"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
