package main

import "fmt"

type SwapAAAAA struct {
	a int
	b int
}

func main() {
	s := SwapAAAAA{
		a: 1,
		b: 2,
	}
	swapStruct(&s)
	fmt.Printf("%+v\n", s)
}

func swapStruct(s *SwapAAAAA) {
	s.a, s.b = s.b, s.a
}
