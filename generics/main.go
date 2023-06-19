package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)


func Add[T constraints.Ordered](a T, b T) T {

	return a + b

}

func main() {
	fmt.Println("Additions is : ", Add(12.2 , 32.4 ))

}
