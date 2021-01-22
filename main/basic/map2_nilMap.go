package main

import "fmt"

func main() {
	var ages map[string]int
	ages1 := make(map[string]int, 0)
	ages2 := map[string]int{}
	ages3 := map[string]int(nil)
	fmt.Println(ages == nil)
	fmt.Println(ages1 == nil)
	fmt.Println(ages2 == nil)
	fmt.Println(ages3 == nil)
}
