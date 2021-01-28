package main

import "fmt"

//構造体
//構造体とポインタ

type Point struct {
	A int
	B string
	C float64
}

func Update(p Point)  {
	p.A = 100
	p.B = "Update"
	p.C = 2.14
}

func Update2(p *Point)  {
	p.A = 100
	p.B = "Update"
	p.C = 2.14
}

func main()  {
	p := Point{}
	Update(p)
	fmt.Println(p)

	//参照渡し
	//アドレスを生成する
	//こちらが推奨されている
	p2 := &Point{}
	Update2(p2)
	fmt.Println(p2)

	//非推奨
	Update2(&p)
	fmt.Println(p)

	//非推奨
	p3 := new(Point)
	Update2(p3)
	fmt.Println(p3)
}