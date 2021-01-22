package main

import "fmt"

type buyer2 interface {
	buy()
}

type student2 struct {
	name string
}

func (s *student2) buy() {
	fmt.Println(s.name, "买买买")
}

func showBuyer2(b buyer2) {
	fmt.Println(b)
}

func main() {
	s := student2{"小明"}
	(&s).buy()
	s.buy()

	//showBuyer2(s) //Cannot use 's' (type student2) as type buyer2Type does not implement 'buyer2' as 'buy' method has a pointer receiver
	showBuyer2(&s)
}
