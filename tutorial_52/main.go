package main

//スコープ
import (
	"fmt"

	//別名指定
	//f"fmt"
	//省略して書ける(非推奨)
	//."fmt"
	//本来は絶対パスで指定
	"./foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	//var b string = s
	b = s
	{
		b := "BBBB"
		fmt.Println(b)
	}
	return b
}

func main() {
	fmt.Println(foo.Max)
	//fmt.Println(foo.min)

	fmt.Println(foo.ReturnMin())

	fmt.Println(appName())
	//fmt.Println(AppName, Version)

	fmt.Println(Do("AAA"))
}
