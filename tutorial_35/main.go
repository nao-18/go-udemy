package main

import "fmt"

//map

func main()  {
	var m = map[string]int{"A": 100, "B": 300}
	fmt.Println(m)

	m2 := map[string]int{"N": 879, "K": 723982}
	fmt.Println(m2)

	//改行する場合は巻末にコロンをつける
	m3 := map[int]string{
		1: "A",
		2: "I",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	m4[0] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	//取得
	fmt.Println(m["A"])
	fmt.Println(m4[2])
	//ない物を取得する場合は初期値が入る
	fmt.Println(m4[3])
	//そのためmap取得はエラーハンドリングがついている
	s, ok := m4[0]
	fmt.Println(s, ok)
	i, result := m4[10]
	if !result {
		fmt.Println(i, "error")
	}

	//値の更新
	m4[0] = "US"
	fmt.Println(m4)

	//追加
	m4[3] = "JJJJJ"
	fmt.Println(m4)

	//削除
	delete(m4, 3)
	fmt.Println(m4)

	//要素数
	fmt.Println(len(m4))
}