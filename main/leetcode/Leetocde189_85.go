package main

import "fmt"

func main() {
	nums := []int{1}
	rotate(nums, 0)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	reverse(nums, 0, len(nums)-1)

	reverse(nums, 0, k%len(nums)-1)

	reverse(nums, k%len(nums), len(nums)-1)

}

func reverse(nums []int, st int, end int) {
	if end <= st {
		return
	}
	for i := 0; i <= (end-st)/2; i++ {
		nums[st+i], nums[end-i] = nums[end-i], nums[st+i]
	}
}
