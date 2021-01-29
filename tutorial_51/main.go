package main

import "fmt"

//interface
//fmp.Stringer

/*
type Stringer interface {
	String() string
}
*/

type Point struct {
	A int
	B string
}

//表示を変える
func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
}
