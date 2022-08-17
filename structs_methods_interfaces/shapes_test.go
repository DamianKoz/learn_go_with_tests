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

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got %v, Want %v", got, want)
		}
	}

	t.Run("rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
