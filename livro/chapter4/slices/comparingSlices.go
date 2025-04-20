// Comparing Slices
// Slices aren't comparable
package main

import (
	"fmt"
)

func main() {
	a := []string{"test1", "test2", "test3", "test4"}
	b := []string{"alpha", "beta", "gama"}
	c := []string{"test1", "test2", "test3", "test4"}

	fmt.Printf("%t\n", equal(a, b))
	fmt.Printf("%t\n", equal(a, c))
	fmt.Printf("%t\n", equal(b, c))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
