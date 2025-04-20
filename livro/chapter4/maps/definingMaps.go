// Hash table | Dispersion table
// Map = hash table reference
package main

import (
	"fmt"
)

func main(){
	ages := make(map[string]int)
	fmt.Println(ages)

	ages2 := map[string]int{
		"alice": 31,
		"bruno": 34,
	}
	fmt.Println(ages2)

	ages3 := make(map[string]int)
	ages3["alice"] = 31
	ages3["charlie"] = 34
	fmt.Printf("alice %v\n", ages3["alice"])
	fmt.Printf("charlie %v\n", ages3["charlie"])

	//deleting item
	//delete(map, key)
	delete(ages3, "charlie")
	fmt.Println(ages3)

	//operations
	ages2["bruno"] = ages2["bruno"] + 1
	fmt.Printf("bruno changed: %v\n", ages2["bruno"])

	ages3["charlie"] += 10
	fmt.Printf("charlie changed: %v\n", ages3["charlie"])

	ages3["alice"]++
	fmt.Printf("alice changed: %v\n", ages3["alice"])

	//Panic
	//Maps element aren't variables, we can't get address
	//_ = &ages2["bruno"]
}