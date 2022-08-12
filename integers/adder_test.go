package integers

import "testing"

func testAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Got: %d, Want: %d")
	}
}
