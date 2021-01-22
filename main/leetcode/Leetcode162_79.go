package main

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if isPeak(nums, mid) {
			return mid
		} else if isPeak(nums, left) {
			return left
		} else if isPeak(nums, right) {
			return right
		}

		if nums[mid] < nums[mid-1] {
			right = mid
		} else {
			left = mid
		}
		left++
		right--
	}
	return -1
}

func isPeak(nums []int, i int) bool {
	if len(nums) <= 1 {
		return i >= 0 && i <= len(nums)-1
	}
	if i == 0 {
		return nums[i] > nums[i+1]
	}
	if i == len(nums)-1 {
		return nums[i] > nums[i-1]
	}
	return nums[i] > nums[i-1] && nums[i] > nums[i+1]
}
