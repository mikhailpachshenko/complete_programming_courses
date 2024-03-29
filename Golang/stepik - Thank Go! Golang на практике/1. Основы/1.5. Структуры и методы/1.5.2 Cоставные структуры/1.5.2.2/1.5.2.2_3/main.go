package main

import (
	"fmt"
)

type person struct {
	firstName  string
	secondName string
}

type book struct {
	title  string
	author person
}

type pBook struct {
	title  string
	author *person
}

func main() {
	a := person{
		firstName:  "Mikhail",
		secondName: "Pachshenko",
	}

	someBook := book{
		title:  "На меньшее я не согласен",
		author: a,
	}

	someBook.author.firstName = "Igor"
	fmt.Println(someBook, someBook.author)
	fmt.Println(a)

	b := &person{"Kirill", "Khivincev"}
	bBook := pBook{"Скорость", b}
	bBook.author.firstName = "Nikita"
	fmt.Println(b)
	fmt.Println(bBook.title, *bBook.author)
}
