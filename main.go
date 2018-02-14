package main

import "fmt"

func main() {
	//test how init func works
	fmt.Println("#### testing how init func works")
	SelectNumber()

	//test Unitialized valiable
	fmt.Println("#### testing Unitialized valiable")
	PrintUninitializedPerson()

	//test method sets
	fmt.Println("#### testing method sets")
	admin1 := admin{"Nick"}

	//The following will fail since a value is passed and not a pointer
	//sendNotification(admin1)
	SendNotification(&admin1)

}
