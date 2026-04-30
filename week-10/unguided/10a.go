package main

import "fmt"

func main() {
	var N, i, j int
	var kelinci [1000]float64
	var min, max float64

	fmt.Print("input N: ")
	fmt.Scan(&N)

	for i = 0; i < N; i++ {
		fmt.Printf("input berat ke-%d: ", i+1)
		fmt.Scan(&kelinci[i])
	}

	min = kelinci[0]
	max = kelinci[0]

	for j = 1; j < N; j++ {
		if kelinci[j] < min {
			min = kelinci[j]
		}

		if kelinci[j] > max {
			max = kelinci[j]
		}
	}

	fmt.Println(min, max)
}
