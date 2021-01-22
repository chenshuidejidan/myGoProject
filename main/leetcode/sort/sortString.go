package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 8, 5, 5, 3, 8}
	s := sort.IntSlice(nums)
	sort.Sort(sort.Reverse(s))
	fmt.Println(s) // [1 3 5 5 8 8]

	strs := []string{"acd", "ab", "d", "b"}
	sort.Strings(strs)
	fmt.Println(strs) // [ab acd b d]
}
