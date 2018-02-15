package entities

type vehicle struct{
	Age int
}

// Car is exported outside of the entities package,
// but vehicle is not
type Car struct{
	vehicle
	NumberPlate int
}