package iteration

import (
	"fmt"
	"testing"
)

func TestRepet(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func ExampleRepeat() {
	rep := Repeat("a")
	fmt.Println(rep)
	// Output: aaaaa
}
