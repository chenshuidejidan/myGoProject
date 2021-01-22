package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count_str := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	var i int
	for i < 5 && input.Scan() {
		count_str[input.Text()]++
		i++
	}

	for line, count := range count_str {
		if count > 1 {
			fmt.Printf("%q\t%d\n", line, count)
		}
	}
}
