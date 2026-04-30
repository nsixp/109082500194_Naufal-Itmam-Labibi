# <h1 align="center">Laporan Praktikum Modul 10 - Pencarian Nilai Max/Min</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 10A]
#### 10a.go

```go
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
```

##### Output 
![Screenshot Output Unguided 10A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-10/unguided/output/output-10a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai N sebagai jumlah anak kelinci. Kemudian, menginputkan sebanyak N kali berat masing-masing kelinci yg akan disimpan ke dalam array. Lalu, program menggunakan variabel min dan max yg diisi dgn nilai pertama array kelinci dan mencari array dari index 1 sampai N-1 untuk membandingkan setiap elemen. Jika lebih kecil dari min maka min diperbarui, jika lebih besar dari max maka max diperbarui. Setelah itu, program akan menampilkan output berat kelinci terkecil dan terbesar.

### 2. [Soal 10B]
#### 10b.go

```go
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
```

##### Output 
![Screenshot Output Unguided 10B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-10/unguided/output/output-10b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai x (jumlah ikan) dan y (kapasitas per wadah). Lalu, menginputkan sebanyak x kali berat masing-masing ikan yg disimpan ke dalam array ikan. Kemudian, program menghitung jumlah wadah agar sisa ikan yg tidak memenuhi satu wadah tetap mendapat wadah sendiri. Setelah itu, program menampilkan outpu total berat tiap wadah, rata-rata berat ikan per wadah.

### 3. [Soal 10C]
#### 10c.go

```go
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
```

##### Output 
![Screenshot Output Unguided 10c](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-10/unguided/output/output-10c.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai n (jumlah balita). Lalu, menginputkan sebanyak n kali berat masing-masing balita yg disimpan mulai dari index 1, sedangkan index 0 digunakan untuk menyimpan nilai n itu sendiri. Kemudian, program membaca jumlah data dari arrBerat[0]. Lalu, mencari array dari index 1 sampai n untuk mencari nilai terkecil dan terbesar, dan menjumlahkan seluruh berat dari index 1 sampai n kemudian membaginya dgn n dan mengembalikan hasilnya. Setelah itu, program menampilkan output berat minimum, maximum, dan rata-rata balita.