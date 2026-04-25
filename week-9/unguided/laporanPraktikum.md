# <h1 align="center">Laporan Praktikum Modul 9 - Array</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 9A]
#### 9a.go

```go
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) < float64(c.r)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik
	var dl1, dl2 bool

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&t.x, &t.y)

	dl1 = didalam(l1, t)
	dl2 = didalam(l2, t)

	if dl1 && dl2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dl1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dl2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
```

##### Output 
![Screenshot Output Unguided 5A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-9/unguided/output/output-9a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan 3 baris bilangan. Baris pertama dan kedua terdiri dari titik koordinat pusat serta radius lingkaran (a, b, c, dan d) dan baris ketiga adalah titik koordinat yg akan diperiksa. Kemudian, program akan menghitung jarak antara titik (a, b) dan (c, d) dgn menggunakan rumus jarak Euclidean. Lalu, program akan menentukan posisi suatu titik tersebut berada di dalam atau di luar lingkaran, jika jarak < radius, maka titik berada di dalam lingkaran. Setelah itu, program akan menampilkan satu output dari empat kemungkinan, yaitu: "Titik di dalam lingkaran 1 dan 2" (jika di kedua lingkaran), "Titik di dalam lingkaran 1" (jika hanya di lingkaran 1), "Titik di dalam lingkaran 2" (jika hanya di lingkaran 2), atau "Titik di luar lingkaran 1 dan 2" (jika tidak di keduanya).

### 2. [Soal 9B]
#### 9b.go

```go
package main

import (
	"fmt"
	"math"
)

func showOddIndex(arr []int) {
	var i int

	for i = 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
}

func showEvenIndex(arr []int) {
	var i int

	for i = 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
}

func showMultipleIndex(arr []int) {
	var x, i int

	fmt.Print("input x: ")
	fmt.Scan(&x)

	for i = 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
}

func deleteIndex(arr []int) []int {
	var index int

	fmt.Print("input index array: ")
	fmt.Scan(&index)

	return append(arr[:index], arr[index+1:]...)
}

func showMean(arr []int) float64 {
	var r, i int

	for i = 0; i < len(arr); i++ {
		r += arr[i]
	}
	return float64(r) / float64(len(arr))
}

func showStdDev(arr []int) float64 {
	var i int
	var m, s, jk float64

	m = showMean(arr)

	for i = 0; i < len(arr); i++ {
		s = float64(arr[i]) - m
		jk += s * s
	}

	return math.Sqrt(jk / float64(len(arr)))
}

func showFrequency(arr []int) int {
	var c, i, fq int

	fmt.Print("input c: ")
	fmt.Scan(&c)

	for i = 0; i < len(arr); i++ {
		if arr[i] == c {
			fq++
		}
	}

	return fq
}

func main() {
	var n, i int
	var arr []int

	fmt.Print("input jumlah array: ")
	fmt.Scan(&n)

	arr = make([]int, n)

	for i = 0; i < n; i++ {
		fmt.Printf("input elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println()
	fmt.Println("a. Menampilkan keseluruhan isi dari array.")
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("b. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
	showOddIndex(arr)

	fmt.Println()
	fmt.Println("c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).")
	showEvenIndex(arr)

	fmt.Println()
	fmt.Println("d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.")
	showMultipleIndex(arr)

	fmt.Println()
	fmt.Println("e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil")
	arr = deleteIndex(arr)
	fmt.Println(arr)

	fmt.Println()
	fmt.Println("f. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
	fmt.Println(showMean(arr))

	fmt.Println()
	fmt.Println("g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.")
	fmt.Println(showStdDev(arr))

	fmt.Println()
	fmt.Println("h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.")
	fmt.Println(showFrequency(arr))
}
```

##### Output 
![Screenshot Output Unguided 5B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-9/unguided/output/output-9b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan jumlah elemen array terlebih dahulu, lalu menginputkan setiap elemen array sebanyak jumlah tersebut. Kemudian, program akan menampilkan keseluruhan isi array. Lalu, program akan menampilkan elemen-elemen array dgn indeks ganjil saja. Setelah itu, program akan menampilkan elemen-elemen array dgn indeks genap saja. Selanjutnya, program akan meminta input nilai x dari pengguna, lalu menampilkan elemen-elemen array dgn indeks kelipatan x. Kemudian, program akan meminta input indeks yg akan dihapus, lalu menghapus elemen pada indeks tersebut dan menampilkan array setelah penghapusan. Lalu, program akan menghitung dan menampilkan nilai rata-rata dari seluruh elemen array. Setelah itu, program akan menghitung dan menampilkan standar deviasi dari seluruh elemen array. Terakhir, program akan meminta input sebuah bilangan c, lalu menghitung dan menampilkan frekuensi kemunculan bilangan c tersebut di dalam array.

### 3. [Soal 9C]
#### 9c.go

```go
package main

import (
	"fmt"
)

func main() {
	var ka, kb, p string
	var hasil []string
	var i int
	var pk int

	pk = 1

	fmt.Print("Klub A: ")
	fmt.Scan(&ka)
	fmt.Print("Klub B: ")
	fmt.Scan(&kb)

	for {
		var sa, sb int

		fmt.Printf("Pertandingan %d : ", pk)
		fmt.Scan(&sa, &sb)

		if sa < 0 || sb < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		if sa > sb {
			fmt.Printf("Hasil %d : %s\n", pk, ka)
			hasil = append(hasil, ka)
		} else if sb > sa {
			fmt.Printf("Hasil %d : %s\n", pk, kb)
			hasil = append(hasil, kb)
		} else {
			fmt.Printf("Hasil %d : Draw\n", pk)
		}

		pk++
	}

	fmt.Println("\nDaftar klub yg menang:")
	for i, p = range hasil {
		if p != "" {
			fmt.Printf("%d. %s\n", i+1, p)
		}
	}
}
```

##### Output 
![Screenshot Output Unguided 5c](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-5/unguided/output/output-9c.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nama Klub A dan Klub B terlebih dahulu. Kemudian, program akan memasuki perulangan untuk setiap pertandingan. Pada setiap pertandingan, program akan meminta input skor Klub A dan Klub B. Jika skor yg dimasukkan bernilai negatif, program akan menghentikan perulangan dan menampilkan pesan "Pertandingan selesai". Jika tidak, program akan membandingkan skor kedua klub: jika skor Klub A lebih besar, program akan menampilkan pemenangnya adalah Klub A dan menyimpannya ke dalam daftar hasil; jika skor Klub B lebih besar, program akan menampilkan pemenangnya adalah Klub B dan menyimpannya ke dalam daftar hasil; jika skor kedua klub sama, program akan menampilkan "Draw" tanpa menyimpan pemenang ke daftar. Setelah perulangan berhenti, program akan menampilkan seluruh daftar klub yg pernah menang beserta nomor urut pertandingannya.

### 4. [Soal 9D]
#### 9d.go

```go
package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var k rune

	*n = 0

	for *n < NMAX {
		fmt.Scanf("%c", &k)

		if k == '.' {
			break
		}

		if k != '\n' && k != ' ' {
			t[*n] = k
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}

	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	var i int

	for i = 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	var temp tabel
	var i int

	for i = 0; i < n; i++ {
		temp[i] = t[i]
	}

	balikkanArray(&temp, n)

	for i = 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}

	return true
}

func main() {
	var tab, asli tabel
	var m, i int

	fmt.Print("teks: ")
	isiArray(&tab, &m)

	for i = 0; i < m; i++ {
		asli[i] = tab[i]
	}

	balikkanArray(&tab, m)

	fmt.Print("Reverse Teks: ")
	cetakArray(tab, m)

	fmt.Print("Palindrome? ")
	if palindrome(asli, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
```

##### Output 
![Screenshot Output Unguided 5c](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-5/unguided/output/output-9d.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan sebuah kalimat dalam bentuk rangkaian huruf yg diakhiri dgn tanda titik. Kemudian, program akan membaca setiap karakter yg dimasukkan, mengabaikan karakter spasi dan baris baru, dan menyimpannya ke dalam array tab hingga mencapai batas maksimum 127 karakter atau sampai menemukan tanda titik. Lalu, program akan menyalin isi array tab ke array asli untuk menyimpan teks asli. Kemudian, program akan membalikkan urutan karakter di dalam array tab dgn cara menukar posisi karakter pertama dgn terakhir, kedua dgn kedua terakhir, dan seterusnya. Setelah itu, program akan menampilkan isi dari array tab yg sudah dibalik. Selanjutnya, program akan mengecek apakah teks asli yg tersimpan di array asli merupakan palindrome dgn cara membandingkan setiap karakter pada teks asli dgn teks yg sudah dibalik. Jika seluruh karakternya sama, maka teks tersebut adalah palindrome. Setelah itu, program akan menampilkan true jika teks tersebut adalah palindrome atau false jika bukan.