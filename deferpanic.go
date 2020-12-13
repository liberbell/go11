package main

func main()  {
	defer func() {
		str := recover()
		fmt.Println(str)
	}
}