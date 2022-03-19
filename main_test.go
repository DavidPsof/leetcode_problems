package main

import "testing"

func TestStrStr(t *testing.T) {
	got := StrStr("hello", "ll")
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSearchInsert(t *testing.T) {
	var tests = []struct {
		a   []int
		x   int
		exp int
	}{
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 2, 3, 5}, 6, 4},
		{[]int{1, 2, 3, 7}, 4, 3},
	}

	for i := range tests {
		res := SearchInsert(tests[i].a, tests[i].x)
		if res != tests[i].exp {
			t.Errorf("got %q, wanted %q", res, tests[i].exp)
		}
	}
}
