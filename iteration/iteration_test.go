package iteration

import "testing"

func TestIteration(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"
	if got != expected {
		t.Errorf("Got %q, Want %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
 