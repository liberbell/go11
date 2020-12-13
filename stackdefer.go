package main

import "fmt"

func main() {
	fmt.Println("Counting...")

	for x := 0; x < 5; x++ {
		defer fmt.Println(x)
	}

	fmt.Println("Finished.")
}
