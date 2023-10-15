package main

import (
	"fmt"
)

func main() {
	JKRowling := NewAuthor("JKRowling", "Harry Potter", 52, 100000, 8)
	printAuthor(JKRowling)
	harryPotter := NewBook("J.K. Rowling", 500, 20, 1000000, 19.99, 1997, 9788700631625, true)
	println("before price change")
	printBookInfo(harryPotter)
	println("After price change")
	println()
	increasePrice(harryPotter, 12.99)
	printBookInfo(harryPotter)
	fmt.Println("End of main method")
}
