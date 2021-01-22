package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// 按照 Person.Age 从大到小排序
type PersonSlice []Person

func (a PersonSlice) Len() int           { return len(a) }
func (a PersonSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PersonSlice) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	a := []int{1, 2, 3, 4, 5, 8, 7, 6}
	sort.Ints(a)
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	sort.Sort(sort.IntSlice(a))

	people := []Person{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	//等价于：sort.Sort(sort.Interface(PersonSlice(people)))
	sort.Sort(PersonSlice(people)) // 按照 Age 升序排序
	fmt.Println(people)

	sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 降序排序
	fmt.Println(people)

}
