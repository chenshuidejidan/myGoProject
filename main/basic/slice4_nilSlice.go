package main

import "fmt"

func main() {
	var slice []int
	slice1 := make([]int, 0)
	slice2 := []int{}
	slice3 := []int(nil)
	fmt.Println(slice == nil)
	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)
	fmt.Println(slice3 == nil)
}
