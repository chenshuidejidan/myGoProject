package main

import "fmt"

var a2b = make(chan struct{})
var b2c = make(chan struct{})
var c2a = make(chan struct{})
var ch = make(chan struct{})

func main() {
	go printA()
	go printB()
	go printC()
	c2a <- struct{}{}
	ch <- struct{}{}
}

func printA() {
	for range c2a {
		fmt.Println("A")
		a2b <- struct{}{}
	}
}

func printB() {
	for range a2b {
		fmt.Println("B")
		b2c <- struct{}{}
	}
}

func printC() {
	count := 0
	for range b2c {
		fmt.Println("C")
		count++
		if count >= 3 {
			close(c2a)
			break
		}
		c2a <- struct{}{}
	}
	<-ch
}
