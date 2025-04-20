package main

import "fmt"

func main() {
	fmt.Println("start")

	var a [3]int
	var q [3]int = [...]int{1, 2, 3} //use "..." to set length of array
	var r [3]int = [3]int{1, 2}

	//test error array
	//
	//var z [3]int = [4]int{1, 2, 3, 4}

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("------")
	for _, v := range q {
		fmt.Println(v)
	}

	fmt.Println("------")
	for _, v := range r {
		fmt.Println(v)
	}
}
