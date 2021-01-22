package main

import "fmt"

func main() {
	s := "你好123"
	fmt.Println(s[0])
	for i, c := range s {
		fmt.Println(i, c)
	}
}
