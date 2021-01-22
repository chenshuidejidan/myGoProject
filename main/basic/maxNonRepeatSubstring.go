package main

import (
	"fmt"
	"time"
)

func main() {
	var str string
	fmt.Println("============最大不重复的连续子串===========")
	for {
		fmt.Print("-> ")
		fmt.Scanf("%s", &str)
		start := time.Now()
		var len int
		var starts []int
		for i := 0; i < 100; i++ {
			len, starts = maxSubstring(str)
		}
		fmt.Println("执行用时：", time.Since(start))
		fmt.Println("最大不重复子串长度为：", len, ",开始位置为：", starts)
	}

}

func maxSubstring(s string) (int, []int) {
	if len(s) < 1 {
		return 0, make([]int, 0)
	}
	nearest := make(map[string]int)
	left, max := 0, 0
	starts := []int{}
	for right, c := range s {
		m := string(c)
		if near, ok := nearest[m]; ok && near >= left {
			left = near + 1
		}
		nearest[m] = right
		if right-left+1 > max {
			max = right - left + 1
			starts = []int{}
			starts = append(starts, 0)
		} else if right-left+1 == max {
			starts = append(starts, left)
		}
	}
	return max, starts
}
