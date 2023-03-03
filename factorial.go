package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("%d! = %d\n", n, factorial(n))
}
