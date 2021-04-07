package main

import "testing"

func AssertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q\n", got, want)
	}
}
