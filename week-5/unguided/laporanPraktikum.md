# <h1 align="center">Laporan Praktikum Modul 5 - Rekursif</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 5A]
#### 5a.go

```go
package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n, i int

	fmt.Print("input n: ")
	fmt.Scan(&n)

	fmt.Print("n: ")
	for i = 1; i <= n; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	fmt.Print("S(n): ")
	for i = 1; i <= n; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
```

##### Output 
![Screenshot Output Unguided 5A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-5/unguided/output/output-5a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan bilangan bulat n. Kemudian, program akan mencari bilangan-bilangan fibonacci dari bilangan n dgn menggunakan subprogram rekursif yg didalamnya juga menggunakan rumus S(n) = S(n-1) + S(n-2). Setelah itu, program akan menampilkan output deret bilangan n dan deret bilangan fibonaccinya.

### 2. [Soal 5B]
#### 5b.go

```go
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
```

##### Output 
![Screenshot Output Unguided 5B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-5/unguided/output/output-5b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan bilangan n untuk jumlah bintang dan baris yg ingin kita tampilkan. Kemudian, program akan mencetak pola bintang dari 1 sampai n. Setiap nilai i bertambah 1, subprogram rekursif pola akan membuat baris baru hingga nilai i menyentuh n. Selain itu, subprogram rekursif bintang juga menambah jumlah bintang yg ditampilkan dalam 1 baris hingga nilai i menyentuh 1. Setelah itu, program akan menampilkan output hasil pola bintangnya.

### 3. [Soal 5C]
#### 5c.go

```go
package main

import "fmt"

func faktor(n, i int) {
	if i > n {
		return
	}

	if n%i == 0 {
		fmt.Printf("%d ", i)
	}

	faktor(n, i+1)
}

func main() {
	var n int

	fmt.Print("input n: ")
	fmt.Scan(&n)

	faktor(n, 1)
}
```

##### Output 
![Screenshot Output Unguided 5c](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-5/unguided/output/output-5c.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan bilangan bulat positif n. Kemudian, program akan mencari bilangan-bilangan faktor dari n dgn menggunakan subprogram rekursif faktor. Subprogram mencari setiap bilangan faktor dgn mengecek apabila sisa bagi n dgn i adalah 0, maka nilai i akan ditampilkan hingga nilai i menyentuh nilai n. Setelah itu, program akan menampilkan output deret bilangan faktor dari n. 