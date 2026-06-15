package main

import "fmt"

type arrInt [2023]int

func insertionSort(T *arrInt, n int) {
	var temp, i, j int

	i = 1
	for i <= n-1 {
		j = i
		temp = T[i]

		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}

		T[j] = temp
		i = i + 1
	}
}

func cekJarak(T arrInt, n int) {
	var jarak, i int
	var tetap bool

	if n < 2 {
		fmt.Println("Data berjarak 0")
	} else {
		jarak = T[1] - T[0]
		tetap = true

		i = 2
		for i <= n-1 {
			if T[i]-T[i-1] != jarak {
				tetap = false
			}
			i = i + 1
		}

		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}

func tampilArray(T arrInt, n int) {
	var i int

	i = 0
	for i <= n-1 {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
		i = i + 1
	}
	fmt.Println()
}

func main() {
	var data arrInt
	var x, n int

	n = 0

	fmt.Scan(&x)
	for x >= 0 {
		data[n] = x
		n = n + 1
		fmt.Scan(&x)
	}

	insertionSort(&data, n)

	tampilArray(data, n)
	cekJarak(data, n)
}
