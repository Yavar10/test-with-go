package main

import "testing"

func TestSplit(t *testing.T) {
	got1, got2 := split(50)
	want1, want2 := 22, 28
	if got1 != want1 && got2 != want2 {
		t.Errorf("got1 %d want1 %d got2 %d want %d", got1, want1, got2, want2)
	}
}
