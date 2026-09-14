package main

import (
	"fmt"
	"os"

	"github.com/goldengeek22/learning-go-programming/ch03-greeter/greeting"
)

func main() {
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	fmt.Println(greeting.Greet(name))
}
