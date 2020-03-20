package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	excepted := "aaaaa"

	if repeated != excepted {
		t.Errorf("excepted %q but got %q", excepted, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeater := Repeat("b", 4)
	fmt.Println(repeater)
	//output bbbb
}
