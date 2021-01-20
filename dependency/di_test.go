package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Fernando")

	got := buffer.String()
	want := "Hello, Fernando"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
