// Reversing a integer slice
package main

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
