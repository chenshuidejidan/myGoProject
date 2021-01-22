package main

import "fmt"

func main() {
	c := make(chan int /*, 1*/)
	fmt.Println(len(c)) //0
	c <- 1
	fmt.Println(<-c) //fatal error: all goroutines are asleep - deadlock!
}
