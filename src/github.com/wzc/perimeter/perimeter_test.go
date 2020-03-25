package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 1.0}
	got := Perimeter(rectangle)
	want := 22.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("want %g but got %g", want, got)
	// 	}
	// }

	// t.Run("ractangles", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 2.0}
	// 	want := 20.0
	// 	checkArea(t, rectangle, want)

	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

	//table driven tests
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
