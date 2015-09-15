package rabinkarp_test

import (
	"fmt"

	"github.com/glkz/rabinkarp"
)

func Example() {
	txt := "a man a plan a canal panama"
	patterns := []string{"man", "boat", "plan", "ana", "banana"}
	matches := rabinkarp.Search(txt, patterns)

	fmt.Println(matches)

	// Output:
	// [man plan ana]
}
