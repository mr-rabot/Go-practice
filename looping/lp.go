package main

import "fmt"


func main()  {


	defer fmt.Println("\n thank you for usei \n")
	for i := 1; i < 10 ; i++ {
		if i==5 {
			fmt.Println("half : ",i)
			
		}else{

			fmt.Println(i)

		}
	}
}