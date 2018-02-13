package main

import "fmt"

type person struct {
	name string
	age int
	male bool
}

func printUninitializedPerson(){
	var p1 person
	fmt.Println(p1)
}
