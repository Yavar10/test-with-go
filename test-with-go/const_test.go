package main

import (
	"math"
	"testing"
)

func TestNeedInt(t *testing.T) {
	got := needInt(Small)
	want := 21

	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestNeedFloatWithSmall(t *testing.T) {
	got := needFloat(Small)
	want := 0.2

	if got != want {
		t.Fatalf("got %f, want %f", got, want)
	}
}

func TestNeedFloatWithBig(t *testing.T) {
	got := needFloat(Big)
	want := float64(Big) * 0.1

	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("got %f, want %f", got, want)
	}
}