package main

import "fmt"

func main() {

	defer fmt.Println("everyone")
	defer fmt.Println("World")

	fmt.Println("Hello")
}
