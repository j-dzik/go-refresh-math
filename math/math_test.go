package math

import "testing"

func TestDouble(t *testing.T) {
	want := 10
	if got := Double(5); got != want {
		t.Errorf("Double(5) == %d want %d", got, want)
	}
}

func TestSquare(t *testing.T) {
	want := 25
	if got := Square(5); got != want {
		t.Errorf("Square(5) == %d want %d ", got, want)
	}
}
