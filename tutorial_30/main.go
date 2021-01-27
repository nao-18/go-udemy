package main

import "fmt"

//スライス
//宣言、操作

func main()  {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{1000, 2000}
	fmt.Println(sl2)

	sl3 := []string{"A", "HKJS"}
	fmt.Println(sl3)

	//初期要素数を指定
	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 555
	fmt.Println(sl2)

	sl5 := []int{1, 4, 4, 8}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	fmt.Println(sl5[2:4])

	fmt.Println(sl5[:2])

	fmt.Println(sl5[:2])

	//最後のみ取得
	fmt.Println(sl5[len(sl5)-1])
	//最初と最後以外全て取得する
	fmt.Println(sl5[1: len(sl5)-1])
}