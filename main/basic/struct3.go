package main

import "fmt"

type m struct {
	name string
	age  int
}

type b struct {
	m
	name string
	age  int
}

func main() {
	s := b{m{"mname", 1}, "bname", 2}
	fmt.Println(s.m.name)
}
