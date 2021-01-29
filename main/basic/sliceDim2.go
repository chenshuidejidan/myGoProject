package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("sliceDim2.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	slice := [][]int{{1, 2}, {2, 3, 4}, {5}}
	for i, j := range slice {
		fmt.Println(i, "\t", j)
	}
}
