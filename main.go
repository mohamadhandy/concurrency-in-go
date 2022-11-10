package main

import "fmt"

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("One!")
	printSomething("Two!")
}
