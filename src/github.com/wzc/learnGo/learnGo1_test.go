package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("wzc")
	want := "hello world,wzc"
	if got != want {
		t.Errorf("got %q wat %q", got, want)
	}
}
