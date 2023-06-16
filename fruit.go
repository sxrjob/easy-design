package main

import "fmt"

type Fruit interface {
	Show()
}

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("i am apple")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("i am banana")
}

type Factory struct {
}
