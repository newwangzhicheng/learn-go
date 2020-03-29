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
	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "121212"
		dictionary := Dictionary{}
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "123456"
		newDefinition := "654321"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("exist word", func(t *testing.T) {
		word := "test"
		definition := "123456"
		dictionary := Dictionary{word: definition}
		newDefinition := "654321"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "123456"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrorWordNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	difinition := "123456"
	dictionary := Dictionary{word: difinition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrorNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()

	if err != want {
		t.Errorf("want %q got %q", want, err)
	}
	if err == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

//验证dictionary中word和difinition是否一致
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added world:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
