package main

import "fmt"

//型
//浮動小数点数

func main()  {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	// 基本的に使用しない
	var fl32 float32 = 3.0
	fmt.Println(fl32)
	fmt.Printf("%T\n", fl32)

	// 暗黙的な宣言はfloat64
	flInit := 3.8
	fmt.Println(fl64 + flInit)
	fmt.Printf("%T, %T\n", fl64, flInit)

	//cast
	fmt.Printf("%T\n", float64(fl32))

	zero := 0.0
	// プラスの無限
	pinf := 1.0 / zero
	fmt.Println(pinf)

	// マイナスの無限
	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}