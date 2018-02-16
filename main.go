package main

import (
	"fmt"
	"github.com/Go-Playground/entities"
	"sync"
)

const(
	one = 1 + iota
	two
	tree
	six = 6
	seven
)

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

	//testing the visibility of unexported & exported types
	//in other packages
	fmt.Println("#### testing package visibility")
	//var myCar Car
	var myCar entities.Car

	//Only NumberPlate is exported from package, this does not mean that
	//we do not have access to properties of vehicle
	myCar = entities.Car{
		NumberPlate: 1234,
	}
	//Age that is exposed value of the unexposed vehicle can still be accessed
	//it is promoted to Car.
	myCar.Age = 5
	fmt.Println(myCar)

	//Test unbuffered channels
	fmt.Println("#### testing unbuffered channels")
	var wg sync.WaitGroup
	wg.Add(2)
	score := make(chan int)
	go Player("Nick", score, &wg)
	go Player("Marg", score, &wg)
	score <- 1
	wg.Wait()

	//Test iota
	fmt.Println("#### testing iota")
	fmt.Printf("%v %v %v %v %v", one, two, tree, six, seven)

}
