package main

import "fmt"

func main() {
	x := 4
	fmt.Println("This is ", x)
	switch x {
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}
}
