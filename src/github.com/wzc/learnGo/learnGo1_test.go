package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello world to people", func(t *testing.T) {
		got := Hello("wzc", "")
		want := "hello world,wzc"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when supply an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world,wsq"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "hola world,Elodie"
		assertCorrectMessage(t, got, want)
	})
}
