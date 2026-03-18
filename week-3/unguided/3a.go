package main

import "fmt"

func faktorial(n int) int {
	var hasil int = 1

	for i := 1; i <= n; i++ {
		hasil *= i
	}

	return hasil
}

func permutasi(n, r int) int {
	if n < r {
		return 0
	}

	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	if n < r {
		return 0
	}

	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	var p1, p2, k1, k2 int

	fmt.Scan(&a, &b, &c, &d)

	p1 = permutasi(a, c)
	k1 = kombinasi(a, c)
	p2 = permutasi(b, d)
	k2 = kombinasi(b, d)

	fmt.Println(p1, k1)
	fmt.Println(p2, k2)
}
