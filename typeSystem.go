package main

import "fmt"

type person struct {
	name string
	age  int
	male bool
}

func PrintUninitializedPerson() {
	var p1 person
	fmt.Println(p1)
}
