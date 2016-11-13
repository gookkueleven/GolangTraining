package main

import "fmt"
var glob_var int = 10
func main() {
	printUTF8()
}

func printUTF8(){
	for i := 60; i < 123; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \t \n", i, i, i, i)
	}
}






