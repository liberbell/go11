package main

import "fmt"

func multi(a, b string) (string, string) {
	return a, b
}

func main() {
	x, y := multi("hello", "world")
	fmt.Println(x, y)
}
