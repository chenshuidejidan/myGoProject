package main

import (
	"fmt"
	"sync"
)

type Student3 struct {
	name string
}
type Student4 struct {
	name string
}

func (s Student3) GetName() string { return s.name }
func (s Student4) GetName() string { return s.name }

type GetNamer interface {
	GetName() string
}

func NewStudent3() *Student3 {
	return &Student3{}
}

func NewStudent4() *Student4 {
	return &Student4{}
}

var mp map[int]GetNamer
var once sync.Once

func newStudents() map[int]GetNamer {
	once.Do(func() {
		mp = make(map[int]GetNamer)
		mp[1] = NewStudent3()
		mp[2] = NewStudent4()
	})
	return mp
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for range make([]struct{}, 10) {
		go func() {
			fmt.Println(newStudents())
			wg.Done()
		}()
	}
	var m Student3
	fmt.Println(m.GetName())
	wg.Wait()
}
