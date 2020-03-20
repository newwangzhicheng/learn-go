package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expect '%d' but got '%d'", expected, sum)
	}
}

//Example will be added to godec
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}
