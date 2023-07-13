package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr := 0, 1 // Initialize the previous and current Fibonacci numbers

	return func() int {
		result := prev
		prev, curr = curr, prev+curr // Calculate the next Fibonacci number
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
