package main

import "fmt"

func bintang(i int) {
	if i > 0 {
		fmt.Print("*")
		bintang(i - 1)
	}
}

func pola(i, n int) {
	if i <= n {
		bintang(i)
		fmt.Println()
		pola(i+1, n)
	}
}

func main() {
	var n int

	fmt.Print("input n: ")
	fmt.Scan(&n)

	pola(1, n)
}
