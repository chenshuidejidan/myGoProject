package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"

	i := 1
	v := reflect.ValueOf(&i) //获取变量指针的反射对象
	v.Elem().SetInt(10)      //获取指针指向的变量,修改其值
	fmt.Println(i)           //10

}
