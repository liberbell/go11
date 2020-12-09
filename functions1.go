package main

import "fmt"

func fun(y int) {
	y = 0
}

func main() {
	y := 10
	fun(y)

	fmt.Println(y)
}
