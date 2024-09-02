package itertools_test

import (
	"fmt"

	it "github.com/astonm/go-itertools"
)

func ExampleNewSeq() {
	for v := range it.NewSeq(1, 2, 3) {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleFromSlice() {
	for v := range it.FromSlice([]int{1, 2, 3}) {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleEnumerate() {
	for i, v := range it.Enumerate(it.NewSeq("a", "b", "c")) {
		fmt.Println(i, v)
	}
	// Output:
	// 0 a
	// 1 b
	// 2 c
}
