package main

import "fmt"

type buyer interface {
	buy()
}

type student struct {
	name string
}

func (s student) buy() {
	fmt.Println(s.name, "买买买")
}

func showBuyer(b buyer) {
	fmt.Println(b)
}

func main() {
	s := student{"小明"}
	s.buy()
	(&s).buy()

	showBuyer(s)
	showBuyer(&s)
}
