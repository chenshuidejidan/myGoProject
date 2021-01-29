package main

import "fmt"

func main() {
	i := functionReturn()(1, 2)
	fmt.Println(i)
}

func functionReturn() func(int, int) int {
	return func(a int, b int) int {
		return a + b
	}
}
