package main

import "fmt"

func main() {
	//a := [5]int{0, 1, 2, 3, 4}
	//b := a[1:]
	//c := a[1:2]
	//c = append(c, 200)
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)

	a := [5]int{0, 1, 2, 3, 4}
	b := a[1:]
	c := a[1:2:2]
	c = append(c, 200)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
