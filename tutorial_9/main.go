package main

import (
	"fmt"
	"strconv"
)

//型変換

func main()  {
	//int →　float64
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)

	//float64 → int
	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)

	//float64 → int(小数点以下切り捨て)
	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	// _(アンダーバー)にすることで破棄できる
	i4, _ := strconv.Atoi(s)
	fmt.Println(i4)
	fmt.Printf("i4 = %T\n", i4)

	//本来の使い方
	i5, err := strconv.Atoi("FDSAHKJFHD")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i5)
	fmt.Printf("i5 = %T\n", i5)

	//数値から文字列へ
	var i6 int = 200
	s6 := strconv.Itoa(i6)
	fmt.Println(s6)
	fmt.Printf("s6 = %T\n", s6)

	//文字列　バイト配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}