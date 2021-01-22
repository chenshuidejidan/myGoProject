package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{
		name: "xiaoming",
		age:  10,
	}
	p.addAge()
	fmt.Printf("%v\n", p)
}

func (p *person) addAge() {
	p.age = p.age + 1
}
