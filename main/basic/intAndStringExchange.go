package main

import (
	"fmt"
	"strconv"
)

func main() {
	in := int64(11)
	a := strconv.FormatInt(in, 2)
	fmt.Println(a)
}
