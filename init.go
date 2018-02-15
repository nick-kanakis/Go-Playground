package main

import "fmt"

var number int

//Init will run before main
func init() {
	number = 12
}

//SelectNumber would print a number if init runs before main
func SelectNumber() {
	if number == 0 {
		fmt.Println("Number is not initialized")
		return
	}
	fmt.Printf("number is %v", number)

}
