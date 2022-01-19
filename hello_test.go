package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Ane")
	want := "Hello, Ane"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
