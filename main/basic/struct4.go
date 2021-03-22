package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type T struct {
	Name string
	Age  int
}

func main() {
	t1 := &T{"aaa", 1}
	t2 := &T{"bbb", 2}
	t3 := &T{"ccc", 3}
	list := []*T{t1, t2, t3}
	bytes, _ := json.Marshal(list)
	fmt.Println(bytes)
	fmt.Println(string(bytes))
	fmt.Println("=============")

	TradeDateBegin := "2020-11-02"
	if TradeDateBegin == "" {
		TradeDateBegin = time.Now().AddDate(0, 0, -3).Format("2006-01-02")
	}
	begin, err := time.ParseInLocation("2006-01-02 15:04:05", TradeDateBegin, time.Local)
	if err != nil {
		log.Println("parse fail....")
	}
	TradeDateBegin = begin.Format("2006-01-02 15:04:05")

	fmt.Println(TradeDateBegin)
}
