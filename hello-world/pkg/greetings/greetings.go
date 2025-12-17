package greetings

import "fmt"

func SayHelloCustomized(name string, newLine bool) {
	fmt.Printf("Hello %s", name)
	fmt.Printf("!")
	if newLine {
		fmt.Printf("\n")
	}
}

func SayHello() {
	SayHelloCustomized("world", false)
}
