package main

import "fmt"

//関数
//無名関数

func main()  {
	f := func(x, y int) int {
		return x + y
	}
	i := f(10, 20)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i2)
}