package main

import "fmt"

func main() {
	//var ages map[string]int

	ages := make(map[string]int)

	ages["lilei"] = 1
	ages["xiaohong"] = 2
	for name, age := range ages {
		fmt.Println(name, age)
	}
}
