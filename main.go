package main

import "fmt"

func main() {
	fmt.Println("First Go Program")
	fmt.Println("What the heck is going on?")

	myName := `Ken`
	fmt.Printf("Hi, my name's %v. I just learned to initialize a %T with Go!\n",
		myName, myName)
}
