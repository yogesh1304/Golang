package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println("Hello World...")
	// fmt.Println("Enter Your Name")
	// var Name string
	// fmt.Scanf("%s", &Name)
	// fmt.Printf("Welcome %s", Name)

	fmt.Print("Enter your name: ")

	r := bufio.NewReader(os.Stdin)

	name, err := r.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Welcome",name)
}
