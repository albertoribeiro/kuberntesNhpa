package main

import "testing"

func TestSqrt(t *testing.T) {
	got := sqrt(5)
	want := 25.00

	if got != want {
		t.Errorf("greeting('TEST') \n got: %v \n want: %v", got, want)
	}
}

func TestGreeting(t *testing.T) {
	got := greeting("TEST")
	want := "<b>TEST</b>"

	if got != want {
		t.Errorf("greeting('TEST') \n got: %v \n want: %v", got, want)
	}
}
