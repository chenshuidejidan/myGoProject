package main

import (
	"fmt"
	"time"
)

func main() {
	local, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-30 18:16:15", local)
	//fmt.Println(t)
	t = t.AddDate(0, 1, 0)
	fmt.Println(t)
}
