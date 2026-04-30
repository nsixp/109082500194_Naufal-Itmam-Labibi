package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, i, jmlWadah, wadahKe, ikanDiWadah int
	var rata2 float64
	var ikan [1000]float64
	var totalPerWadah [1000]float64

	fmt.Print("input x & y: ")
	fmt.Scan(&x, &y)

	for i = 0; i < x; i++ {
		fmt.Printf("input berat ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	jmlWadah = int(math.Ceil(float64(x) / float64(y)))

	for i = 0; i < x; i++ {
		wadahKe = i / y
		totalPerWadah[wadahKe] += ikan[i]
	}

	for i = 0; i < jmlWadah; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(totalPerWadah[i])
	}

	fmt.Println()

	for i = 0; i < jmlWadah; i++ {
		ikanDiWadah = y
		if i == jmlWadah-1 {
			ikanDiWadah = x - (i * y)
		}

		rata2 = totalPerWadah[i] / float64(ikanDiWadah)

		if i > 0 {
			fmt.Print(" ")
		}

		fmt.Print(rata2)
	}

	fmt.Println()
}
