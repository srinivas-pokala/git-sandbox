package main_test

import (
	"testing"

	main "."
)

func TestGreet(t *testing.T) {
	if main.Greet() != "Hello, World!" {
		t.Fail()
	}
}
