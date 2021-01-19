package main

import "fmt"

func main() {

	// 明示的な宣言
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go!"
	fmt.Println(s)

	var b bool = true
	fmt.Println(b)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 200
	s3 = "Golang!"
	fmt.Println(i3, s3)

	i = 400
	fmt.Println(i)

	// 暗黙的な宣言
	// 変数名 := 値
	i4 := 5000
	fmt.Println(i4)

	i4 += 400
	fmt.Println(i4)

	//i4 = "hello!"
	//fmt.Println(i4)


}
