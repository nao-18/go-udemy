package main

import "fmt"

//channel
//複数のゴルーチン間でのデータ受け渡しをする為に設計されたデータ構造。
//キュー(先入先出)
//宣言、操作

func main() {
	//宣言
	//双方向
	var ch1 chan int
	//受信専用
	//var ch2 <- chan int
	//送信専用
	//var ch3 -> chan int

	ch1 = make(chan int)

	ch2 := make(chan int)
	//バッファサイズを調べる
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))
	//バッファサイズを指定できる
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	//送信
	ch3 <- 1
	fmt.Println(len(ch3))
	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	//受信
	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))
	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	/*
		バッファを超えた場合,
		デットロックになる。
	*/
	//ch3 <- 1
	//ch3 <- 2
	//ch3 <- 3
	//ch3 <- 4
	//ch3 <- 5
	//ch3 <- 6
}
