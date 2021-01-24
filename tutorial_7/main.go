package main

import "fmt"

//配列型
//後から要素数を変更できない

func main()  {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//要素の数を自動で入れてくれる(暗黙的宣言)
	arr4 := [...]string{"C"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2[2])

	//配列内の型が同じでも要素数が違えば、同一データ型だと認識されない。
	//var arr5 [4]int
	//arr5 = arr1
	//fmt.Println(arr5)
}