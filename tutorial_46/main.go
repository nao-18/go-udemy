package main

import "fmt"

//構造体
//埋め込み→クラス継承のようなもの

type Point struct {
	A int
	B string
	C float64
}

type BigPoint struct {
	//省略して記入可能
	//Point Point
	Point
}

func (p *Point) Set(i int) {
	p.A = i
}

func main()  {
	bp := BigPoint{}
	fmt.Println(bp)
	//pointを更新
	bp.Point.A = 100
	bp.Point.B = "Apple"
	bp.Point.C = 2.9
	fmt.Println(bp)
	//省略した場合は直接アクセス可能
	fmt.Println(bp.A)

	//初期値の設定
	bp2 := BigPoint{
		Point: Point{
			A: 100,
			B: "Banana",
			C: 2.978,
		},
	}
	fmt.Println(bp2)
	bp2.Set(2938)
	fmt.Println(bp2)

}