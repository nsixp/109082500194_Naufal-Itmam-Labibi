package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	var n, i int

	n = int(arrBerat[0])
	*bMin = arrBerat[1]
	*bMax = arrBerat[1]

	for i = 1; i <= n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	var n, i int
	var total float64

	n = int(arrBerat[0])

	for i = 1; i <= n; i++ {
		total += arrBerat[i]
	}

	return total / float64(n)
}

func main() {
	var balita arrBalita
	var n, i int
	var bMin, bMax, rata2 float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	balita[0] = float64(n)

	for i = 1; i <= n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i)
		fmt.Scan(&balita[i])
	}

	hitungMinMax(balita, &bMin, &bMax)
	rata2 = rerata(balita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata2)
}
