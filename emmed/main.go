package main

import (
	_ "embed"
	"fmt"
	"strings"
)

var (
	//go:embed colors.txt
	data []byte
)

func main() {

	fmt.Println("--------------------")

	lines := strings.Split(string(data), "\r\n")

	var colours string
	for _, line := range lines {
		if line != "" {
			val := line
			colours = val

		}
	}
	fmt.Println("Printing  : ", colours)

}
