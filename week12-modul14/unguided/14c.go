package main

import (
	"fmt"
)

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

func median(T *arrInt, n int) int {
	if n == 0 {
		return 0
	}

	if n%2 == 1 {
		return T[n/2]
	} else {
		return (T[n/2-1] + T[n/2]) / 2
	}
}

func main() {
	var T arrInt
	var x, count int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			if count > 0 {
				selectionSort(&T, count)
				hasil := median(&T, count)
				fmt.Println(hasil)
			}
		} else {
			T[count] = x
			count++
		}
	}
}
