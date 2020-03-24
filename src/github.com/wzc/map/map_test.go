package mapdic

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "123456"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "123456"

		assertString(t, want, got)
	})

	t.Run("unknow word", func(t *testing.T) {
		_, err := dictionary.Search("unknow")

		assertError(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "121212")

	got, err := dictionary.Search("test")
	want := "121212"

	if err != nil {
		assertError(t, err, ErrorNotFound)
	}
	assertString(t, want, got)
}

func assertError(t *testing.T, err, want error) {
	t.Helper()

	if err != want {
		t.Errorf("want %q got %q", want, err)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}
