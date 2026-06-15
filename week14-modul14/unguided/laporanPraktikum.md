# <h1 align="center">Laporan Praktikum Modul 14 - Insertion Sort</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 14A]
#### 14a.go

```go
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
```

##### Output 
![Screenshot Output Unguided 14A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week14-modul14/unguided/output/output-14a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan beberapa bilangan bulat yg diakhiri dgn bilangan negatif. Bilangan positif akan disimpan ke dalam array, sedangkan bilangan negatif tidak disimpan. Setelah itu, program mengurutkan data secara ascending, yaitu dengan mengambil satu elemen sebagai temp, lalu membandingkannya dengan elemen-elemen sebelumnya. Jika elemen sebelumnya lebih besar dari temp, maka elemen tersebut digeser ke kanan sampai ditemukan posisi yang tepat. Setelah array terurut, program menampilkan semua isi array. Kemudian program mengecek apakah jarak antar data selalu sama dengan membandingkan selisih setiap elemen yang berurutan. Jika semua selisih sama, program menampilkan "Data berjarak X", sedangkan jika ada selisih yang berbeda, program menampilkan "Data berjarak tidak tetap".

### 2. [Soal 14B]
#### 14b.go

```go
package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax + 1]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	var i int

	fmt.Scan(n)

	i = 1
	for i <= *n {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)

		i = i + 1
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i, idxMax int

	idxMax = 1

	i = 2
	for i <= n {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
		i = i + 1
	}

	fmt.Println(
		pustaka[idxMax].judul,
		pustaka[idxMax].penulis,
		pustaka[idxMax].penerbit,
		pustaka[idxMax].tahun,
	)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var i, j int

	i = 2
	for i <= n {
		j = i
		temp = pustaka[i]

		for j > 1 && temp.rating > pustaka[j-1].rating {
			pustaka[j] = pustaka[j-1]
			j = j - 1
		}

		pustaka[j] = temp
		i = i + 1
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	if n < 5 {
		batas = n
	} else {
		batas = 5
	}

	i = 1
	for i <= batas {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(pustaka[i].judul)
		i = i + 1
	}
	fmt.Println()
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	var kiri, kanan, tengah int
	var ketemu bool

	kiri = 1
	kanan = n
	ketemu = false

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			ketemu = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu {
		fmt.Println(
			pustaka[tengah].judul,
			pustaka[tengah].penulis,
			pustaka[tengah].penerbit,
			pustaka[tengah].tahun,
			pustaka[tengah].eksemplar,
			pustaka[tengah].rating,
		)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var Pustaka DaftarBuku
	var nPustaka int
	var r int

	DaftarkanBuku(&Pustaka, &nPustaka)

	CetakTerfavorit(Pustaka, nPustaka)

	UrutBuku(&Pustaka, nPustaka)

	Cetak5Terbaru(Pustaka, nPustaka)

	fmt.Scan(&r)
	CariBuku(Pustaka, nPustaka, r)
}
```

##### Output 
![Screenshot Output Unguided 14B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week14-modul14/unguided/output/output-14b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai N sebagai jumlah buku. Setelah itu, kita menginputkan data setiap buku. Kemudian, program mencari buku terfavorit dgn membandingkan rating setiap buku dan memilih buku dgn rating tertinggi. Setelah itu, program mengurutkan data buku berdasarkan rating secara descending dgn mengambil satu data buku sebagai temp, lalu menggeser data sebelumnya ke kanan jika rating data sebelumnya lebih kecil dari rating temp. Setelah itu, program menampilkan maksimal lima judul buku dengan rating tertinggi. Terakhir, program membaca nilai rating yang ingin dicari, kemudian melakukan binary search pada data yang sudah terurut. Terakhir, program menampilkan data buku dgn rating tersebut.