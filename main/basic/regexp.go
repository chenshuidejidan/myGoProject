package main

import (
	"fmt"
)

var aa string
var done bool

func setup() {
	aa = "hello"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	fmt.Println(aa)
}
