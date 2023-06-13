package main

import "fmt"

func main() {

	var ob Animal = Animals{
		noLegs:   4,
		Name:     "Dog",
		Sound:    "Bow BOw",
		isEating: true,
	}

	var ob1 Animal = Animals{
		noLegs:   4,
		Name:     "Cat",
		Sound:    "Mew Mew",
		isEating: true,
	}
	fmt.Println("\n")
	ob.Details()
	ob.showName()
	fmt.Println("\n--------------------\n")
	ob1.Details()

}

type Animal interface {
	Details()
	showName() string
}

type Animals struct {
	Name     string
	noLegs   int
	Sound    string
	isEating bool
}

func (animal Animals) Details() {

	fmt.Println("Name of Animal : ", animal.Name)
	fmt.Println("number of legs : ", animal.noLegs)
	fmt.Println("Sound : ", animal.Sound)
	fmt.Println("Is Eating : ", animal.isEating)
}

func (animal Animals) showName() string {

	return animal.Name

}
