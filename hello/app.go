package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	fmt.Print("\nEnter Your Name : ")

	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	name = strings.ToUpper(name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hii mr." + name + " Welcome to code \n")

}
