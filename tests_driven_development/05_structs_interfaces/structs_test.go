package main

import "testing"

// test for Perimeter function
func TestPermimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but want %.2f", got, want)
	}
}

// test for Area function
// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
// 	t.Run("for rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		want := 72.0
// 		checkArea(t, rectangle, want)
// 	})
// 	t.Run("for circles", func(t *testing.T) {
// 		circle := Circle{10} // πr2
// 		want := 314.1592653589793
// 		checkArea(t, circle, want)
// 	})
// }

// table driven tests
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			shape:   Rectangle{12, 6},
			hasArea: 72.0,
		},
		{
			shape:   Circle{10},
			hasArea: 314.1592653589793,
		},
		{
			shape:   Triangle{12, 6},
			hasArea: 16.0,
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt, got, tt.hasArea)
			}
		})
	}
}
