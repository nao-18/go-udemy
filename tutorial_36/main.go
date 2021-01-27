package main

import "fmt"

//マップ
//for

func main() {
	n := map[string]int{
		"Apple":  100,
		"Banana": 200,
	}

	for k, v := range n {
		fmt.Println(k, v)
	}
}
