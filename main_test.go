package main

import "testing"

func Test_greet(t *testing.T) {
	got := greet()
	want := "Hello!"
	if got != want {
		t.Errorf("greet() = %v, want %v", got, want)
	}
}
