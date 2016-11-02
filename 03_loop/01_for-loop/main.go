package main

import "fmt"

func main() {
	for i := 60; i < 123; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \t \n", i, i, i, i)
	}
}
