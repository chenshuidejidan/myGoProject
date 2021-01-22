package main

import "fmt"

func main() {
	s := []string{"a", "b"}
	s = append(s, "c", "d", "e")
	fmt.Printf("%d %d", len(s), cap(s))

	a := []int{1, 2}
	a = append(a, 4, 5, 6)
	fmt.Printf("%d %d", len(a), cap(a))
}
