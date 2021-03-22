package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	a := int64(1)
	b := 1
	fmt.Println(a == 1)
	fmt.Println(b == 1)
	//fmt.Println(a == b)   // mismatched types int64 and int

	c := 1
	fmt.Println(reflect.TypeOf(c))

	d := 3 * 0.33
	fmt.Println(reflect.TypeOf(d))

	fmt.Println(a == math.MaxInt64)
	//fmt.Println(numSquares(12))
}

func numSquares(n int) int {
	if n == 0 {
		return 1
	}
	squares := generateSquares(n)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		min := i
		for j := 1; j < len(squares) && i-squares[j] >= 0; j++ {
			if dp[i-squares[j]]+1 < min {
				min = dp[i-squares[j]] + 1
			}
		}
		dp[i] = min
	}
	return dp[n]
}

func generateSquares(n int) []int {
	res := make([]int, 0, 101)
	for i := 0; i*i <= n && i <= 100; i++ {
		res = append(res, i*i)
	}
	return res
}
