package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {

	fmt.Printf("Hello World!\n")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	// fmt.Println(puppy.From12())

	puppy.From12()

	fmt.Printf("This is version v1.6")
}

//version v1.6
//This is a test
