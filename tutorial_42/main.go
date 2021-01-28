package main

import "fmt"

//ポインタ
//アドレスと型情報
func Double(i int) {
	i = i * 2
}

//ポインタ型を引数に取る場合
func Doublev2(i *int) {
	*i = *i * 2
}

//スライスの場合
func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	//int等は値型と呼ばれる
	var n int = 100
	fmt.Println(n)
	//ポインタを出力
	fmt.Println(&n)

	Double(n)
	fmt.Println(n)

	//ポインタ型の宣言
	var p *int = &n
	fmt.Println(p)
	//実態の値を表示(デリファレンス)
	fmt.Println(*p)

	/*
		*p = 300
		fmt.Println(n)
		n = 2000
		fmt.Println(*p)
	*/

	Doublev2(&n)
	fmt.Println(n)
	Doublev2(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)
}
