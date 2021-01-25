package main

import "fmt"

//関数
//ジェネレーター
//何らかのルールに従って連続した値を返し続ける仕組み

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main()  {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	ints2 := integers()
	fmt.Println(ints2())
	fmt.Println(ints2())
	fmt.Println(ints2())
	fmt.Println(ints2())
}