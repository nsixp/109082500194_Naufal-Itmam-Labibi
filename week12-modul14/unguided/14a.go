package main

import "fmt"

type arrInt [1000000]int

func selectionSort(T *arrInt, n int) {
	var t, i, j, idx_min int

	i = 1
	for i <= n-1 {
		idx_min = i - 1
		j = i
		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i = i + 1
	}
}

func main() {
	var n, m, i, j int
	var rumah arrInt

	fmt.Scan(&n)

	i = 0
	for i < n {
		fmt.Scan(&m)

		j = 0
		for j < m {
			fmt.Scan(&rumah[j])
			j = j + 1
		}

		selectionSort(&rumah, m)

		j = 0
		for j < m {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
			j = j + 1
		}
		fmt.Println()

		i = i + 1
	}
}
