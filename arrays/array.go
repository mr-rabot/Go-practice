package main

import "fmt"

func main() {

	var arr1 = [...]int64{1, 2, 3, 4, 5, 6}
	arr2 := [...]string{"jaya", "gogi", "soni"}
	fmt.Println("\n-----------------------")
	fmt.Println("\nfirst array is : ", arr1)

	e := 0
	for _, n := range arr1 {

		fmt.Printf("\n Element %v : %v", e, n)
		e++
	}
	fmt.Println("\n-----------------------")
	fmt.Println("\nsecond array is : ", arr2)
	el := 0
	for _, n := range arr2 {

		fmt.Printf("\n Element %v : %v", el, n)
		el++
	}
	fmt.Println("\n thank you")

}
