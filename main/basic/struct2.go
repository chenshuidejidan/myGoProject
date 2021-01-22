package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

type color string

func (p point) distance(q point) float64 {
	fmt.Println("point's distance() is runing")
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

type coleredPoint struct {
	point
	color
}

func (cp coleredPoint) distance(cp2 coleredPoint) float64 {
	fmt.Println("cp's distance() is runing...")
	return math.Sqrt(math.Pow(float64(cp.x-cp2.x), 2) + math.Pow(float64(cp.y-cp2.y), 2))
}

func (cp coleredPoint) show() {
	fmt.Println("coloredPoint.show()........")
}

func (c color) distance(c2 color) float64 {
	fmt.Println("color's distance() is runing  ", c, "   ", c2)
	return 0.0
}

func main() {
	cp := coleredPoint{
		point{1, 1},
		"red",
	}
	cp2 := coleredPoint{point{2, 2}, "green"}

	cpdistance := coleredPoint.distance
	pdistance := cp.point.distance
	cdistance := cp.color.distance

	cpshow := coleredPoint.show

	cpshow(cp)

	fmt.Println(cpdistance(cp, cp2)) //func (cp coleredPoint) distance(cp2 coleredPoint) float64{...}

	fmt.Println(pdistance(cp2.point))
	fmt.Println(cdistance(cp2.color))
}
