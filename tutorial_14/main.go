package main

import "fmt"

//関数

func Plus(x, y int) int {
	return x + y
}

//引数の型はまとめてできる
//func Plus(x, y int) int {
//	return x + y
//}

//複数の返り値がある場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//戻り値の型に変数を書くと、returnに明示しなくても良い
func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return.")
	return
}

func main() {
	i := Plus(10, 9)
	fmt.Println(i)

	i2, i3 := Div(8, 4)
	fmt.Println(i2, i3)

	//使いたくない場合
	i4, _ := Div(10, 5)
	fmt.Println(i4)

	i5 := Double(2000)
	fmt.Println(i5)

	Noreturn()
}
