package main

import (
	"fmt"
)

func main() {
	a := 10
	defer func() {
		fmt.Println(a)
	}()

	defer func(a int) {
		fmt.Println(a)
	}(a)

	a = 100
}
