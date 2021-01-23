package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	dp := make([]int, 2)
	if len(nums) == 0 {
		return 0
	}
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i&1] = max(dp[(i+1)&1], dp[(i+2)&1]+nums[i])
	}
	return dp[(len(nums)-1)&1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
