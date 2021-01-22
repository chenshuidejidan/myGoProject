package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "left foot"
	t := s
	s += "left foot, right foot"[9:]
	fmt.Println(s)
	fmt.Println(t)
	m := "你好"
	fmt.Println(len(m))
	fmt.Println(utf8.RuneCountInString(m))
}
