package iteration

import "testing"

func TestRepeat(t *testing.T) {
	// 声明并初始化
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("Repeated result was incorrect, got: %s, want: %s.", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}
