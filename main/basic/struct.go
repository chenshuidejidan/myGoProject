package main

import "fmt"

type s struct {
	x int
	y int
}

func (s *s) exChange() {
	fmt.Println("s.x的地址：", &(s.x), "   s.y的地址：", &(s.y))
	(s).x, (s).y = (s).y, (s).x
	fmt.Println("s.x的地址：", &(s.x), "   s.y的地址：", &(s.y))
}

func main() {
	myS := &s{1, 2}
	fmt.Println(myS.x, myS.y)
	(myS).exChange()
	fmt.Println(myS.x, myS.y)
}
