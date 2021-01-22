package main

import "fmt"

func main() {
	time := make([]struct{}, 3)
	for range time {
		fmt.Println("--")
	}
}
