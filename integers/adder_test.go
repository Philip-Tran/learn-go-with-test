package integers

import (
	"fmt"
	"testing"
)

func testAdder(t *testing.T) {
	t.Run("Add two number", func(t *testing.T) {
		sum := Add(5, 2)
		expected := 7

		if sum != expected {
			t.Errorf("sum %d expected %d", sum, expected)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
