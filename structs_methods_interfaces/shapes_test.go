package shapes

import "testing"

func TestRectangle(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	CheckErrorMessage := func(got, want float64) {
		if got != want {
			t.Errorf("Got %v, Want %v", got, want)
		}
	}
	t.Run("Rectangle Perimeter", func(t *testing.T) {
		got := rectangle.Perimeter()
		want := 40.0
		CheckErrorMessage(got, want)

	})

	t.Run("Rectangle Area", func(t *testing.T) {
		got := rectangle.Area()
		want := 100.0
		CheckErrorMessage(got, want)
	})
}

func TestCircle(t *testing.T) {
	circle := Circle{10.0}
	CheckErrorMessage := func(got, want float64) {
		if got != want {
			t.Errorf("Got %v, Want %v", got, want)
		}
	}
	t.Run("Circle Area", func(t *testing.T) {
		got := circle.Area()
		want := 314.1592653589793
		CheckErrorMessage(got, want)
	})
}

// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("Got %v, Want %v", got, want)
// 		}
// 	}

// 	t.Run("rectangle Area", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72)
// 	})

// 	t.Run("Circle Area", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.01},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 3}, want: 151},
	}

	for _, shape := range areaTests {
		t.Run(shape.name, func(t *testing.T) {
			got := shape.shape.Area()
			if got != shape.want {
				t.Errorf("%#v Got %v, Want %v", shape.shape, got, shape.want)
			}
		})

	}

}
