package main

import "fmt"

//構造体

type Point struct {
	//フィールドは先頭を大文字にする
	A int
	B string
	C float64
}

func main()  {
	var p Point
	fmt.Println(p)

	p2 := Point{A: 1, B: "Golang", C: 1.23}
	fmt.Println(p2)
	//フィールドの値を取得する
	fmt.Println(p2.A)

	//上書き
	p2.A = 10
	fmt.Println(p2.A)

	//他の初期化
	p3 := Point{1, "Python", 3.14}
	fmt.Println(p3.A)

	p4 := Point{A: 24}
	fmt.Println(p4)
}