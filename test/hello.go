package test

import (
	"testing"
)

func Hello() string {
	return "Hello, World!"
}
func TestHello(t *testing.T) {
	expected := "Hello, World!"
	result := Hello()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
