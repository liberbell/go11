package main

import "fmt"

func num(yPtr *int) {
	*yPtr = 5
}

func main() {
	yPtr := new(int)
	fmt.Println(yPtr)
	num(yPtr)

	fmt.Println(*yPtr)
}
