package main

import "fmt"

func shiftAnd(s string, t string) (res []int) { // s是目标串，t是模式串
	positions := make(map[int32]uint)
	mask := uint(1) << (len(t) - 1)
	shift, D := uint(1), uint(0)
	for _, c := range t {
		//辅助map记录每个字符出现的位置，如a在T[0],T[3]处出现，positions['a'] = 001001.
		positions[c] |= shift
		shift <<= 1
	}
	for i, c := range s {
		if D = (D<<1 | 1) & positions[c]; D&mask > 0 { //t[0..j]是s[0..i]的后缀，则D[j]=1(从低到高第j位)
			res = append(res, i-len(t)+1)
		}
	}
	return res
}

func main() {
	s := "abcabcabcabcaabc"
	t := "abca"
	res := shiftAnd(s, t)
	fmt.Println(res)
}
