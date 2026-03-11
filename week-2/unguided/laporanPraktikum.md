# <h1 align="center">Laporan Praktikum Modul 2 - Review Algoritma dan Pemrograman 1</h1>
<p align="center">Naufal Itmam Labibi - 109082500194</p>

## Unguided 

### 1. [Soal 2A]
#### 2a.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)

	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	temp = satu
	satu = dua
	dua = tiga
	tiga = temp

	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

****
```

##### Output 
![Screenshot Output Unguided 2A](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-2/unguided/output/output-2a.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai string 3x berturut-turut. Setiap nilai string yg kita inputkan akan disimpan ke dalam variabel "satu", "dua", dan "tiga" sesuai dengan urutan inputan kita. Kemudian, nilai di variabel "satu" diganti dengan nilai di variabel "dua", begitu juga pada variabel "dua". Lalu, nilai di variabel "tiga" akan diganti dengan nilai di variabel "temp" yg dimana masih kosong, dan variabel "temp" akan diisi dengan nilai di variabel "satu". Setelah itu, program akan menampilkan output nilai variabel "satu", "dua", dan "tiga".

### 2. [Soal 2B]
#### 2b.go

```go
package main

import "fmt"

func main() {
	var i, j int
	var w1, w2, w3, w4 string
	var berhasil bool

	for i = 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&w1, &w2, &w3, &w4)

		if w1 == "merah" && w2 == "kuning" && w3 == "hijau" && w4 == "ungu" {
			j++
		}

	}

	berhasil = j == 5
	fmt.Println("BERHASIL:", berhasil)
}

```

##### Output 
![Screenshot Output Unguided 2B](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-2/unguided/output/output-2b.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan 4 urutan warna yg sesuai, yaitu "merah", "kuning", "hijau", dan "ungu" sebanyak 5x. Kemudian, program akan mengecek apakah kelima percobaan tersebut berhasil atau tidak. Percobaan dikatakan berhasil apabila urutan warna dari setiap percobaannya sesuai. Setelah itu, program akan menampilkan output nilai variabel berhasil dalam bentuk boolean true atau false.

### 3. [Soal 2C]
#### 2c.go

```go
package main

import "fmt"

func main() {
	var bp, sb, d1, bk, bj, tb int

	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&bp)

	d1 = bp / 1000
	sb = bp % 1000

	if sb >= 500 && bp <= 10000 {
		bk = sb * 5
	} else if sb < 500 && bp <= 10000 {
		bk = sb * 15
	} else {
		bk = sb * 0
	}

	bj = d1 * 10000
	tb = bj + bk

	fmt.Printf("Detail berat: %d kg + %d gr\n", d1, sb)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", bj, bk)
	fmt.Printf("Total biaya pengiriman: Rp. %d\n", tb)
}
```

##### Output 
![Screenshot Output Unguided 2C](https://github.com/nsixp/109082500194_Naufal-Itmam-Labibi/blob/main/week-2/unguided/output/output-2c.png)

##### Penjelasan
Ketika program dijalankan, kita perlu menginputkan nilai berat parsel (gram). Kemudian, program akan menghitung sisa berat parselnya terlebih dahulu. Lalu, program akan mengecek jika sisa berat tidak kurang dari 500 gram, maka parsel akan dikenakan biaya tambahan sebesar Rp. 5,- per gram-nya, tetapi jika sisa berat kurang dari 500 gram, maka parsel akan dikenakan biaya tambahan sebesar Rp. 15,- per gram-nya, dan jika berat parsel lebih dari 10 kg, maka tidak ada biaya tambahan. Setelah itu, program akan menampilkan output dari seluruh hasil perhitungannya.