package main

import "fmt"

type person struct {
	name string
	age  int
	male bool
}

//PrintUninitializedPerson tests how an unitiliazed struct would print
func PrintUninitializedPerson() {
	var p1 person
	fmt.Println(p1)
}
