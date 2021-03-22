package main

import (
	"fmt"
	"math"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	s := "left foot"
	t := s
	s += "left foot, right foot"[9:]
	fmt.Println(s)
	fmt.Println(t)
	m := "你好"
	fmt.Println(len(m))
	fmt.Println(utf8.RuneCountInString(m))
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)).Len)

	fmt.Printf("\n\n\n\n\n")
	fmt.Printf("%x\n", []byte("世"))
	str := "世"
	for _, c := range []byte(str) {
		fmt.Printf("%x ", c)
	}
	fmt.Println()
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()
	for _, c := range str {
		fmt.Printf("%x  ", c)
	}
	fmt.Println()
	fmt.Printf("%#v\n", []rune("世界"))
	fmt.Printf("%#v\n", string([]rune{'世', '界'}))

	var a float64
	a = math.NaN()
	fmt.Println(a)

}
