package main

import (
	"fmt"

	"example.com/myproject/pkg/calculate"
	"example.com/myproject/pkg/greeting"
	"example.com/myproject/pkg/helloworld"
)

func init() {
	fmt.Println("First init in main")
}

func main() {
	fmt.Println("Main function")

	// Call the SayHello function from the helloworld package
	helloworld.SayHello()
	greeting.SayHello()
	response := calculate.Add(1, 2)
	fmt.Printf("Calculate %d + %d = %d\n", 1, 2, response)
}
