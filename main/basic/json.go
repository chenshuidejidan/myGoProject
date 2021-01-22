package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Cat struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

type CatGroup struct {
	Cats  []Cat  `json:"cats"`
	Owner string `json:"owner"`
}

func main() {
	catsOfXiaoming := `{
		"cats" : [
			{"color":"黄色","name":"小黄猫"},
			{"color":"花色","name":"小花猫"}
		],
		"owner" : "小明"
	}`
	var cg CatGroup

	json.NewDecoder(strings.NewReader(catsOfXiaoming)).Decode(&cg)
	fmt.Printf("%+v", cg)
}
