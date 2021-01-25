package main

import (
	"fmt"
	"strconv"
)

//条件分岐
//エラーハンドリング

func main()  {
	var s string = "A"
	//var s string = "1000"
	//i, _ := strconv.Atoi(s)
	//fmt.Printf("i = %T\n", i)

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)
}