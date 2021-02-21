package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	f1(&s1, 2)
	f2(s2, 2)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
}

func f1(s *[]int, a int) {
	*s = (*s)[a:]
}

func f2(s []int, a int) {
	s = s[a:]
}
