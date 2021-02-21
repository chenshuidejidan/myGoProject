package main

import "fmt"

func main() {
	s := "龥🤩"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	for _, c := range s {
		fmt.Printf("%x ", c)
	}
}
