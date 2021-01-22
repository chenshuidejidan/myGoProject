package main

import "fmt"

func main() {
	var funcs []func()
	for i := 0; i < 5; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i)
		})
	}

	for _, fun := range funcs {
		fun()
	}

}
