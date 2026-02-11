package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jephthah")
	wanted := "Hello Jephthah"
	if got != wanted {
		t.Errorf("got %q wanted %q", got, wanted)
	}
}