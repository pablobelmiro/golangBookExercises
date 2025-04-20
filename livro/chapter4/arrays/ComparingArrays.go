package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2}
	var q [3]int = [...]int{1, 2, 3} //use "..." to set length of array
	var r [3]int = [3]int{1, 2}

	fmt.Println(a == q)
	fmt.Println(q == r)
	fmt.Println(a == r)
}
