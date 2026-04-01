# <h1 align="center">Laporan Praktikum Modul 4 - Prosedur</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 4A]
#### 4a.go

```go
package main

import "fmt"

func factorial(n int, hasil *int) {
	var i int

	*hasil = 1

	for i = 2; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int

	factorial(n, &fn)
	factorial(n-r, &fnr)

	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int

	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)

	*hasil = fn / (fr * fnr)
}

func main() {
	var (
		a, b, c, d, p1, p2, c1, c2 int
	)

	fmt.Print("Input a, b, c, & d: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```

##### Output 
![Screenshot Output Unguided 4A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-4/unguided/output/output-4a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan bilangan a, b, c, dan d dgn syarat a >= c dan b >= d. Kemudian, program akan menghitung nilai permutasi a dgn c dan dilanjut menghitung kombinasi a dgn c, begitu juga dengan bilangan b dgn d. Permutasi dihitung dgn menggunakan faktorial terlebih dahulu di dalam subprogram, yg dimana subprogram faktorial akan dipanggil di subprogram permutasi dan kombinasi untuk dihitung hasilnya. Setelah itu, program akan menampilkan output hasil permutasi 1 (a dgn c), kombinasi 1 (a dgn c), permutasi 2 (b dgn d), dan kombinasi 2 (b dgn d).

### 2. [Soal 4B]
#### 4b.go

```go
package main

import "fmt"

func hitungSkor(soal, skor *int) {
	var waktu, i int

	*soal = 0
	*skor = 0

	for i = 1; i < 9; i++ {
		fmt.Print("Input waktu: ")
		fmt.Scan(&waktu)

		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, maxSoal, minSkor int

	for {
		fmt.Print("\nInput nama: ")
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	fmt.Println("\nPemenang:")
	fmt.Println(pemenang, maxSoal, minSkor)
}
```

##### Output 
![Screenshot Output Unguided 4B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-4/unguided/output/output-4b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nama peserta yg kemudian diikuti dgn durasi (menit) peserta tersebut menyelesaikan soal sebanyak 8 kali. Kemudian, program akan menjumlahkan setiap waktu yg diinputkan sampai menyentuh jumlah maksimal soal (8 soal). Selain itu, program juga menghitung jumlah soal yg diselesaikan oleh peserta hingga waktu selesai (301 menit atau 5 jam 1 menit). Program akan terus berjalan hingga kita menginputkan nama peserta dengan kata "Selesai". Setelah itu, program akan menentukan pemenangnya berdasarkan peserta yg paling cepat menyelesaikan semua soal dan menampilkan hasil outputnya berupa nama peserta, jumlah soal yg diselesaikan, serta jumlah waktu yg dihabiskan.