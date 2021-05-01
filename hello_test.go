package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jill")
	want := "Hello, Jill"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}