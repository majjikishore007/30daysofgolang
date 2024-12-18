package itreations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

// Writing Benchmark tests

func  BenchmarkRepeat(b *testing.B){
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}