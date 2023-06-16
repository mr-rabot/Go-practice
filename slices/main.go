package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {

	fmt.Println("\n\t---------------------")
	fmt.Println("\t     Slice Sorting\n")
	s1 := []int{23, 43, 52, 12, 43}
	fmt.Println("before sort : ", s1)
	sort.Ints(s1)
	fmt.Println("after sort : ", s1)
	fmt.Println("\n")
	fmt.Println("----------- Space triming -----------")
	slice := []byte("  Hello, World!  ")
	trimmed := bytes.TrimSpace(slice)
	fmt.Println(string(trimmed))
	fmt.Println("\n")

	res1 := bytes.Trim([]byte("**#**Welcome this is golang **$**"), "*#$")
    res2 := bytes.Trim([]byte("!!!!lets do sum fun with go@@@@"), "!@")
    res3 := bytes.Trim([]byte("^^golang&&$$"), "^$&")
	
    
    fmt.Printf("trimed Slice:\n")
    fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)
    fmt.Printf("\nSlice 3: %s", res3)
	fmt.Println("\n")

}
