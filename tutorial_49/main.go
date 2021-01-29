package main

import "fmt"

//interface
//最もポピュラーな使い方 異なる型に共通の性質を付与する

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 22},
		&Car{Number: "1234", Model: "QD-123"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}
