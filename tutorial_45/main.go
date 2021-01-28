package main

import "fmt"

//構造体
//メソッド

type Point struct {
	A int
	B string
	C float64
}

func Set(p *Point, i int)  {
	p.A = i
}

func (p *Point) Set1(i int) {
	p.A = i
}

func (p Point) Set2(i int)  {
	p.A = i
}

func main()  {
	p1 := &Point{A: 1}
	//あまり良くない
	Set(p1, 1)
	fmt.Println(p1)

	//推奨
	p1.Set1(3)
	fmt.Println(p1)

	p2 := Point{}
	//値渡しでは更新できない
	p2.Set2(100)
	fmt.Println(p2)

	p2.Set1(11111)
	fmt.Println(p2)
}