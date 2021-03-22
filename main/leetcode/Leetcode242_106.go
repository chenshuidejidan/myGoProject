package main

import (
	"fmt"
	"sort"
)

func isAnagram(s, t string) bool {
	sb, tb := []byte(s), []byte(t)
	if len(sb) != len(tb) {
		return false
	}

	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
	sort.Slice(tb, func(i, j int) bool { return tb[i] < tb[j] })
	for i := 0; i < len(sb); i++ {
		if sb[i] != tb[i] {
			return false
		}
	}
	return true
}

type myrune []rune

func (m myrune) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m myrune) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m myrune) Len() int {
	return len(m)
}

func main() {
	s := "annccsbar"

	r := myrune(s)

	sort.Sort(sort.Reverse(r))
	fmt.Println(string(r))
}
