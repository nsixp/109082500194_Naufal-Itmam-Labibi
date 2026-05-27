# <h1 align="center">Laporan Praktikum Modul 14 - Selection Sort</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 14A]
#### 14a.go

```go
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
```

##### Output 
![Screenshot Output Unguided 14A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week12-modul14/unguided/output/output-14a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai N sebagai jumlah baris yang akan diproses. Kemudian untuk setiap baris, kita menginputkan nilai M yang merupakan banyaknya nomor rumah pada baris tersebut, lalu menginputkan sebanyak M kali nomor rumah yang akan disimpan ke dalam array. Program menggunakan selection sort untuk mengurutkan nomor rumah secara ascending dengan cara mencari nilai minimum dari sisa array yang belum terurut, lalu menukarnya dengan elemen pertama dari sisa array tersebut. Setelah array terurut, program menampilkan semua nomor rumah dalam satu baris dipisahkan spasi, dan setiap baris diakhiri dengan baris baru.

### 2. [Soal 14B]
#### 14b.go

```go
package main

import "fmt"

type arrInt [1000000]int

func selectionSortAsc(T *arrInt, n int) {
	var t, i, j, idx_min int

	for i = 1; i <= n-1; {
		idx_min = i - 1

		for j = i; j < n; {
			if T[idx_min] > T[j] {
				idx_min = j
			}

			j++
		}

		t = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = t
		i++
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var t, i, j, idx_max int

	for i = 1; i <= n-1; {
		idx_max = i - 1

		for j = i; j < n; {
			if T[idx_max] < T[j] {
				idx_max = j
			}

			j++
		}

		t = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = t
		i++
	}
}

func main() {
	var n, i, d, jumlah, angka int
	var ganjil, genap arrInt
	var jGanjil, jGenap int

	fmt.Scan(&n)

	for d = 0; d < n; {
		fmt.Scan(&jumlah)

		jGanjil = 0
		jGenap = 0

		for i = 0; i < jumlah; {
			fmt.Scan(&angka)

			if angka%2 != 0 {
				ganjil[jGanjil] = angka
				jGanjil = jGanjil + 1
			} else {
				genap[jGenap] = angka
				jGenap = jGenap + 1
			}

			i++
		}

		selectionSortAsc(&ganjil, jGanjil)
		selectionSortDesc(&genap, jGenap)

		for i = 0; i < jGanjil; {
			if i > 0 {
				fmt.Print(" ")
			}

			fmt.Print(ganjil[i])
			i++
		}

		if jGanjil > 0 && jGenap > 0 {
			fmt.Print(" ")
		}

		for i = 0; i < jGenap; {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(genap[i])

			i++
		}
		fmt.Println()

		d++
	}
}
```

##### Output 
![Screenshot Output Unguided 14B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week12-modul14/unguided/output/output-14b.png)

##### Penjelasan
Ketika program dijalankan, kita menginputkan n sebagai jumlah baris data. Untuk setiap baris, kita menginputkan jumlah bilangan yang akan dimasukkan, lalu menginputkan bilangan-bilangan tersebut. Program memisahkan bilangan ganjil dan genap ke dalam dua array berbeda. Kemudian, program mengurutkan array ganjil secara ascending (membesar) menggunakan selection sort, dan array genap secara descending (mengecil) menggunakan selection sort. Setelah terurut, program menampilkan seluruh bilangan ganjil diikuti bilangan genap dalam satu baris, dipisahkan spasi. Proses ini diulang sebanyak n baris.

### 3. [Soal 14C]
#### 14c.go

```go
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
```

##### Output 
![Screenshot Output Unguided 14c](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week12-modul14/unguided/output/output-14c.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan rangkaian bilangan bulat secara terus-menerus. Program akan terus membaca input sampai menemukan bilangan -5313 yang menandakan akhir dari masukan. Setiap kali kita menginputkan bilangan, program akan memeriksa bilangannya. Jika bilangan tersebut adalah 0, maka program akan mengurutkan semua data yang sudah tersimpan sebelumnya menggunakan algoritma selection sort secara ascending, kemudian menghitung nilai median dari data yang sudah terurut tersebut lalu mencetaknya ke layar. Jika bilangan yang diinputkan bukan 0 dan bukan -5313, maka bilangan tersebut akan disimpan ke dalam array. Untuk menghitung median, jika jumlah data ganjil, median diambil dari nilai tengah array, sedangkan jika jumlah data genap, median adalah hasil pembagian bulat ke bawah dari penjumlahan dua nilai tengah. Proses ini berulang terus hingga menemukan bilangan -5313 yang mengakhiri program.