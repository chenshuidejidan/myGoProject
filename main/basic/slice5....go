package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	add(&a, 1, 3, 5)
	fmt.Printf("%#v\n", a)
}

func add(a *[]int, b ...int) {
	fmt.Printf("%#v\n", b)
	*a = append(*a, b...)
}
