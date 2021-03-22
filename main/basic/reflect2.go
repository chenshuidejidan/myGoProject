package main

import (
	"fmt"
	"reflect"
)

type mmmm struct {
	m1 string
	m2 int
}

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

	fmt.Println("=====Kind和Name====")
	m := mmmm{}
	typeOfm := reflect.TypeOf(m)
	fmt.Println(typeOfm.Name())
	fmt.Println(typeOfm.Kind())
	fmt.Println(typeOfm.String())
	//fmt.Println(typeOfA.Elem())
	fmt.Println("====指针和Elem=====")
	m2 := &mmmm{}
	typeOfm2 := reflect.TypeOf(m2)
	fmt.Println(typeOfm2.Name())
	fmt.Println(typeOfm2.Kind(), typeOfm2.String(), typeOfm2.Elem())

	typeOfm2Elem := typeOfm2.Elem()
	fmt.Println(typeOfm2Elem.Name(), typeOfm2Elem.Kind())

	fmt.Println("====反射获取结构体成员类型=====")
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	mycat := cat{Name: "mimi", Type: 1}
	typeOfCat := reflect.TypeOf(mycat)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

	fmt.Println(reflect.ValueOf(mycat))

	fmt.Println("=====反射修改结构体成员======")
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
