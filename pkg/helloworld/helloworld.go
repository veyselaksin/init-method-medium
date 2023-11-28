package helloworld

import "fmt"

var (
	name string
)

func init() {
	name = "Veysel"
	println("First init in helloworld package")
}

func SayHello() {
	fmt.Printf("Hello %s\n", name)
}
