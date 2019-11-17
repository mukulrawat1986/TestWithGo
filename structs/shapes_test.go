package main

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}

		got := r.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		c := Circle{
			Radius: 10.0,
		}

		got := c.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
