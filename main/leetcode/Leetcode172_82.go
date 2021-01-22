package main

import "fmt"

func main() {
	fmt.Println(trailingZeroes(100))
}

func trailingZeroes(n int) int {
	fiveCount := 0
	for n > 0 {
		n /= 5
		fiveCount += n // 0-n中，有5的i次方的数作为因子的数的个数
	}
	return fiveCount
}
