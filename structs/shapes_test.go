package main

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got %v, want %v", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		r := Rectangle{
			Width:  10.0,
			Height: 10.0,
		}

		checkArea(t, r, 100.0)
	})

	t.Run("circle", func(t *testing.T) {
		c := Circle{
			Radius: 10.0,
		}

		checkArea(t, c, 314.1592653589793)
	})
}
