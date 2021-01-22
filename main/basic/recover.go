package main

import "fmt"

func errGenFunc(i int) {
	arr := [10]int{}
	defer func() {
		recover()
	}()
	fmt.Println(arr[i])
}

func post() { fmt.Println("post....") }

func main() {
	errGenFunc(100)
	post()
}
