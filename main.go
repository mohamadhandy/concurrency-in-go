package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"theta",
		"zeta",
		"epsilon",
		"eta",
		"pi",
		"delta",
	}
	wg.Add(len(words))
	for index, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", index, word), &wg)
	}
	wg.Wait()

	wg.Add(1)
	printSomething("Two!", &wg)
}
