package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("testing rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10,10}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("testing circle perimeter", func(t *testing.T) {
		circle := Circle{10}
		checkPerimeter(t, circle, 62.83185307179586)
	})

	t.Run("testing triangle perimeter", func(t *testing.T) {
		triangle := Triangle{3,4}
		checkPerimeter(t, triangle, 12.0)
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("testing rectangle area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("testing circle area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTable(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12,6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
