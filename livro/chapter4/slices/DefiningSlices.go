package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, v := range Q2 {
			if s == v {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	endlessSumer := months[:5]
	fmt.Printf("endlessSumer: %q\n", endlessSumer)

	//Panic - out of slice!
	fmt.Println(summer[:20])
}
