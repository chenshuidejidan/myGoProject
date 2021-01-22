package main

import (
	"fmt"
	"myLearnProject/main/important"
)

func main() {
	b := important.B{
		BB: 1,
		//bb: 1,  //Unexported field 'bb' usage in struct literal
		//AA: 1,  //Unknown field 'AA' in struct literal
	}
	b.AA = 2
	fmt.Printf("%v", b) //{{0 2} 0 1}
}
