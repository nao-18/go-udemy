package main

import "log"

//スライス
//for
func main()  {
	sl := []string{"a", "b", "c"}
	log.Println(sl)

	for i, v := range sl {
		log.Println(i, v)
	}

	//valueのみ
	for _, v := range sl {
		log.Println(v)
	}

	//indexのみ
	for i := range sl {
		log.Println(i)
	}

	//古典的for
	for i := 0; i < len(sl); i++ {
		log.Println(sl[i])
	}

}