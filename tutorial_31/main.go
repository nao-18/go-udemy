package main

import "fmt"

//スライス
//append make len cap

func main() {
	sl := []int{100, 3000}
	fmt.Println(sl)

	//append
	//値追加
	sl = append(sl, 5000)
	fmt.Println(sl)
	sl = append(sl, 384, 883782)
	fmt.Println(sl)

	//make
	sl2 := make([]int, 5)
	fmt.Println(sl2)
	//要素数
	fmt.Println(len(sl2))
	//容量
	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 11)
	fmt.Println(sl3)
	//要素数
	fmt.Println(len(sl3))
	//容量
	fmt.Println(cap(sl3))

	/**
	不用意にcap(容量)を設定すると、
	要素数が容量を超えた場合に倍の容量が確保されて余分にメモリを食ってしまう。
	*/
	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sl3)
	//要素数
	fmt.Println(len(sl3))
	//容量
	fmt.Println(cap(sl3))
}
