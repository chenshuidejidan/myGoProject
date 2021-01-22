package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		fmt.Printf("v:%d, v_addr:%p, elem_addr:%p\n", v, &v, &slice[i])
	}
}

// v:1, v_addr:0xc000018080, elem_addr:0xc000014180
// v:2, v_addr:0xc000018080, elem_addr:0xc000014188
// v:3, v_addr:0xc000018080, elem_addr:0xc000014190
