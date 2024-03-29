package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	t.Parallel()

	t.Run("rectangle",
		func(t *testing.T) {
			t.Parallel()

			rect := Rectangle{10.0, 10.0}
			got := rect.Perimeter()
			want := 40.0
			assert.InEpsilon(t, want, got, 1e-4)
		})
	t.Run("circle",
		func(t *testing.T) {
			t.Parallel()

			rect := Circle{10.0}
			got := rect.Perimeter()
			want := 62.8318530718
			assert.InEpsilon(t, want, got, 1e-4)
		})
}

func TestArea(t *testing.T) {
	t.Parallel()

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, testCase := range areaTests {
		testCase := testCase
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got := testCase.shape.Area()
			if got != testCase.hasArea {
				t.Errorf("%#v got %g want %g", testCase.shape, got, testCase.hasArea)
			}
		})
	}
}
