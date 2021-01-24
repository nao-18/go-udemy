package main

import "fmt"

//定数
const Pi = 3.14

const (
	URL = "http://google.com"
	SiteName = "google"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

//定数はオーバーフローしない
//var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

//連続する整数の連番を生成する(iota)
const (
	c0 = iota
	c1
	c2
)

func main()  {
	fmt.Println(Pi)

	//Pi = 3
	//fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	//値がない場合は前の値が入る
	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)
}