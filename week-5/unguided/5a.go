package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n, i int

	fmt.Print("input n: ")
	fmt.Scan(&n)

	fmt.Print("n: ")
	for i = 0; i <= n; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	fmt.Print("S(n): ")
	for i = 0; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
