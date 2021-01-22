package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "a中国"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println()
	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}
	s = "a中🐶"
	fmt.Printf("% x\n", s) // "61 e4 b8 ad f0 9f 90 b6"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[61 4e2d 1f436]"
	fmt.Printf("%T\n", r)
	fmt.Println(string(r)) // "a中🐶"
}
