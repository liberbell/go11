package main

import "fmt"

func main() {
	for y := 1; y <= 10; y++ {
		if y%2 == 0 {
			fmt.Println(y, "even number")
		} else {
			fmt.Println(y, "odd number")
		}
	}
}
