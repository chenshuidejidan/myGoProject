package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{0, 0, 0, 12, 123, 120}
	fmt.Println(largestNumber(nums))
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
