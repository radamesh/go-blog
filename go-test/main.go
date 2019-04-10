package main

import "fmt"

func main()  {
	fmt.Println("1 + 1 = ", 1.0 + 1.0)
	fmt.Println(len("Hello world"))
	fmt.Println("Hello world"[1])
	fmt.Println("Hello " + "world")

	var world string = "Hello World"

	fmt.Println(world[0])

	for _, worlds := range world {
		fmt.Println(worlds)
	}
}
