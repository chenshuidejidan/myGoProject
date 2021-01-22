package main

import "fmt"

const (
	apple int = iota
	banana
	orange
)

func main() {
	arr := []int{apple: 3, orange: 1, banana: 2, 10: 10}
	fmt.Println(arr)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	fmt.Println(a == b) //true

	changeArr(a)
	fmt.Println(a)

	fmt.Println("===============")
	p := &a
	fmt.Println((*p)[1])
	fmt.Println(p[1])
	fmt.Println(&a[0])
	fmt.Println(&a)
	fmt.Printf("%p\n", &a)
	fmt.Println((&a)[1])

	fmt.Println("===============")

}

func changeArr(arr [2]int) {
	arr = [2]int{9, 9}
}
