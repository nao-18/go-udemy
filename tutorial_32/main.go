package main

import "log"

//スライス
//copy

func main()  {
	//そのまま変数に渡すと参照渡しになる
	//sl := []int{100, 200}
	//sl2 := sl
	//sl2[0] = 777
	//log.Println(sl, sl2)

	//値渡しはcopyを使う
	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	log.Println(sl2)

	n := copy(sl2, sl)
	sl2[0] = 88888
	log.Println(n, sl2)
	log.Println(sl)
}