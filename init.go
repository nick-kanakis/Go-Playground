package main

import "fmt"

var number int

//Init will run before main
func init() {
	number = 12
}

func main() {
	if number == 0 {
		fmt.Println("Number is not initialized")
		return
	}
	fmt.Printf("number is %v", number)

}
