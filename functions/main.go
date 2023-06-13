package main

import "fmt"

func main()  {

	fmt.Println(msg())
	var sum = add(23,43)
	fmt.Println("\n addition is :",sum,"\n")

	fmt.Println(summ(2,3,4,5,34,65,76,87,34,34,34))
	fmt.Println("\n")
	
}

func add(v1 , v2 float64) float64 {

	return v1 + v2
	
}

func msg() string  {

	var msg = " \n welcome to my code "

	return msg
	
}

func summ(vals ...int) int  {

	sum :=0
	
	for _, n := range vals{
		sum += n
	}

	return sum
	
}