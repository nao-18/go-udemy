package main

import "fmt"

//interface
//カスタムエラー

/*
type error interface {
	Error() string
}
*/

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 500}
}

func main()  {
	err := RaiseError()
	fmt.Println(err.Error())

	//そのままでは直接アクセスできない
	//fmt.Println(err.Message)

	//型アサーションでアクセスする
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
