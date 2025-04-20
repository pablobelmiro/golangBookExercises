//"zero" value of a slice is nil

package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("%v\n", s)

	s = nil
	fmt.Printf("%v\n", s)

	s = []int(nil)
	fmt.Printf("%v\n", s)

	s = []int{}
	fmt.Printf("%v\n", s)

	fmt.Printf("%t\n", len(s) == 0)
}