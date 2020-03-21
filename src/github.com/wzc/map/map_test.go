package mapDic

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "123456"}

	got := Search(dictionary, "test")
	want := "123456"

	if got != want {
		t.Errorf("got '%s' want '%s', key is '%s'", got, want, "test")
	}
}
