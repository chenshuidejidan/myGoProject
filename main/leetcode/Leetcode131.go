package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	res := make([][]string, 0)
	isPositvie := IsPositvieString(s)
	dfs2(s, 0, make([]string, 0), isPositvie, &res)
	return res
}

func dfs2(s string, l int, cur []string, isPositvie [][]bool, res *[][]string) {
	if l == len(s) {
		temp := make([]string, len(cur))
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}
	for i := l; i < len(s); i++ {
		if isPositvie[l][i] {
			cur = append(cur, s[l:i+1])
			dfs2(s, i+1, cur, isPositvie, res)
			cur = cur[:len(cur)-1]
		}
	}
}

func IsPositvieString(s string) [][]bool {
	res := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = make([]bool, len(s))
	}
	for incr := 0; incr < len(s); incr++ {
		for i := 0; i+incr < len(s); i++ {
			if incr <= 1 {
				res[i][i+incr] = s[i] == s[i+incr]
			} else {
				res[i][i+incr] = res[i+1][i+incr-1] && s[i] == s[i+incr]
			}
		}
	}
	return res
}
