package main

import "fmt"

func titleToNumber(s string) int {
	res := 0
	for _, c := range s {
		res *= 26
		temp := int(c - 'A' + 1)
		res += temp
	}
	return res
}

func main() {
	s := "AA"
	fmt.Printf("%d", titleToNumber(s))
}
