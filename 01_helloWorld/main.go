package main

import (
	"fmt"
)

func main() {
	helloThere("Kook")
}

func helloThere(name string) {
	fmt.Println("Hello, " + name)
}
