package main

import "testing"

func TestStrStr(t *testing.T) {
	got := StrStr("hello", "ll")
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
