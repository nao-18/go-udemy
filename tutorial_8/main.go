package main

import "fmt"

//インターフェース

func main() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	x = 1.888
	fmt.Println(x)

	x = "ADF"
	fmt.Println(x)

	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// 計算を行うことはできない。
	//x = 2
	//fmt.Println(x + 3)
}
