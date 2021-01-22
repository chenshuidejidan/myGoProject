package main

import "fmt"

type Animal struct {
	color string
}

type dog struct {
	Animal
	dogType string
}

func (a Animal) showColor() {
	fmt.Println("我的颜色是：", a.color)
}

func (a *Animal) changeColor(newColor string) {
	a.color = newColor
}

type shower interface {
	showColor()
}

type changer interface {
	changeColor(s string)
}

func showerCanShow(s shower) {
	s.showColor()
}

func changerCanChange(c changer) {
	c.changeColor("绿色")
}

func main() {
	dahuang := dog{
		Animal:  Animal{"黄色"},
		dogType: "中华田园犬",
	}
	showerCanShow(dahuang)
	showerCanShow(&dahuang)
	//changerCanChange(dahuang)  //Cannot use 'dahuang' (type dog) as type changerType does not implement 'changer' as 'changeColor' method has a pointer receiver
	changerCanChange(&dahuang)
	fmt.Println(dahuang.color)
}
