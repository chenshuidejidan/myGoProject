package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	slice1 := a[:]
	slice1[1] = 1000
	slice2 := a[1:3]
	slice1[4] = 100
	fmt.Println(slice1)
	slice2 = append(slice2, 1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println("==========")
	var s []int
	fmt.Printf("声明  %t\n", s == nil)
	s = []int(nil)
	fmt.Printf("s=[]int{nil}  %t\n", s == nil)
	s = []int{}
	fmt.Printf("s=[]int{} %t\n", s == nil)

}
