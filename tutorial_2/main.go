package main

import "fmt"

func main()  {
	var i int = 1200

	// 環境依存
	//int8
	//int16
	//int32
	//int64

	var i2 int64 = 200
	fmt.Println(i)

	// 環境依存と明示的なintは計算できない
	//fmt.Println(i+i2)

	//型を表示する
	fmt.Printf("%T\n", i2)
	//cast
	fmt.Printf("%T\n", int32(i2))
	//cast and plus
	fmt.Println(i + int(i2))
}