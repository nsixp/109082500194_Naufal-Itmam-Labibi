# <h1 align="center">Laporan Praktikum Modul 3 - Fungsi</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 3A]
#### 3a.go

```go
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
```

##### Output 
![Screenshot Output Unguided 3A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-3/unguided/output/output-3a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan bilangan a, b, c, dan d dgn syarat a >= c dan b >= d. Kemudian, program akan menghitung nilai permutasi a dgn c dan dilanjut menghitung kombinasi a dgn c, begitu juga dengan bilangan b dgn d. Permutasi dihitung dgn menggunakan faktorial terlebih dahulu di dalam subprogram, yg dimana subprogram faktorial akan dipanggil di subprogram permutasi dan kombinasi untuk dihitung hasilnya. Setelah itu, program akan menampilkan output hasil permutasi 1 (a dgn c), kombinasi 1 (a dgn c), permutasi 2 (b dgn d), dan kombinasi 2 (b dgn d).

### 2. [Soal 3B]
#### 3b.go

```go
package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
```

##### Output 
![Screenshot Output Unguided 3B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-3/unguided/output/output-3b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan bilangan bulat a, b, dan c. Kemudian, program akan menghitung f(x) = x^2, g(x) = x - 2, dan h(x) = x + 1 dgn nilai x diambil dari a, b, dan c. Lalu, program akan menghitung fungsi komposisi (fogoh)(a), (gohof)(b), dan (hofog)(c). Setelah itu, program akan menampilkan output hasil dari ketiga fungsi komposisi tersebut.

### 3. [Soal 3C]
#### 3c.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) < r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64
	var d1, d2 bool

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	d1 = didalam(cx1, cy1, r1, x, y)
	d2 = didalam(cx2, cy2, r2, x, y)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

##### Output 
![Screenshot Output Unguided 3C](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-3/unguided/output/output-3c.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan 3 baris bilangan. Baris pertama dan kedua terdiri dari titik koordinat pusat serta radius lingkaran (a, b, c, dan d) dan baris ketiga adalah titik koordinat yg akan diperiksa. Kemudian, program akan menghitung jarak antara titik (a, b) dan (c, d) dgn menggunakan rumus jarak Euclidean. Lalu, program akan menentukan posisi suatu titik tersebut berada di dalam atau di luar lingkaran, jika jarak < radius, maka titik berada di dalam lingkaran. Setelah itu, program akan menampilkan satu output dari empat kemungkinan, yaitu: "Titik di dalam lingkaran 1 dan 2" (jika di kedua lingkaran), "Titik di dalam lingkaran 1" (jika hanya di lingkaran 1), "Titik di dalam lingkaran 2" (jika hanya di lingkaran 2), atau "Titik di luar lingkaran 1 dan 2" (jika tidak di keduanya).