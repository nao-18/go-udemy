package main

import "fmt"

//for

func main() {
	//i := 0
	//for {
	//	i++
	//	if i == 3 {
	//		break
	//	}
	//	fmt.Println("loop")
	//}
	//
	//point := 0
	//for point < 10 {
	//	fmt.Println(point)
	//	point++
	//}
	//
	//for i := 0; i < 10; i++ {
	//	if i == 3 {
	//		fmt.Println(i)
	//	}
	//}

	//for i := 0; i < 10; i++ {
	//	if i == 3 {
	//		continue
	//	}
	//	if i == 6 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	//古典的for文
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	arr2 := [3]int{4, 5, 6}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	//valueのみ
	for _, v := range arr2 {
		fmt.Println(v)
	}
	//インデックスのみ
	for v := range arr2 {
		fmt.Println(v)
	}

	stringArr := []string{"Python", "Golang", "PHP"}
	for i, v := range stringArr {
		fmt.Println(i, v)
	}

	m := map[string]int{"apple": 1000, "banana": 2000}
	for i, v := range m {
		fmt.Println(i, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}
}
