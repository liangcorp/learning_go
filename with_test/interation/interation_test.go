package interation

import (
	"fmt"
	"testing"
)

func Repeat(charactor string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated = repeated + charactor
	}

	return repeated
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
