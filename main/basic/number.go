package main

import "fmt"

type a []int

func main() {

	aa := a{1, 2, 3}
	aa = append(aa, 4)

	fmt.Println(aa)
	fmt.Printf("%T\n", aa)

	aa.getMax()

}

func (a a) getMax() {
	fmt.Println("a's max")
}
