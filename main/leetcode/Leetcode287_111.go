package main

func main() {
	nums := []int{1, 3, 4, 2, 2}
	findDuplicate(nums)
}

func findDuplicate(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	for i := 1; i <= len(nums); i++ {
		sum -= i
		if sum%(len(nums)-i) == 0 && sum/(len(nums)-i) <= i {
			return i
		}
	}
	return -1
}
