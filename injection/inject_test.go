package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		t.Helper()
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
