package main

import "fmt"

type Student struct {
	StID     int64
	StName   string
	StCourse string
	StYear   string
	StMoblie []string
}

func main() {

	var sTd = Student{
		101,
		"raman",
		"mca",
		"2nd year",
		[]string{"9912670535", ""},
	}

	var st2 Student

	st2.StID=102
	st2.StName="yogesh turerao"
	st2.StCourse="mca"
	st2.StYear="final year"
	st2.StMoblie=[]string{"898989898"}

	fmt.Println("\n\t students ")
	fmt.Println("---------------------------------")

	fmt.Println(sTd)
	fmt.Println(st2)
	fmt.Println("\n")

}
