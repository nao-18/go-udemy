package main

import (
	"fmt"
	"time"
)

//channel
//close

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println("END")
}

func main() {
	ch1 := make(chan int, 2)
	/*
		ch1 <- 1
		close(ch1)

		//closeしたチャネルには送信できない
		//ch1 <- 1

		//が、受信はできる
		//fmt.Println(<-ch1)


		//チャネルのオープン状態を返す。
		//厳密にはチャネルがclose かつ バッファが0の場合にfalseが返る。

		i, ok := <-ch1
		fmt.Println(i, ok)

		i2, ok := <-ch1
		fmt.Println(i2, ok)
	*/

	//チャネルcloseとゴルーチンの実例
	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Millisecond)
}
