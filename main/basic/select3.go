package main

import "fmt"

func main() {
	ch := randomOneZero()
	for i := range ch {
		fmt.Println(i)
	}
}

func randomOneZero() chan int {
	ch := make(chan int)
	go func() {
		for range make([]struct{}, 10) {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
		close(ch)
	}()
	return ch
}
