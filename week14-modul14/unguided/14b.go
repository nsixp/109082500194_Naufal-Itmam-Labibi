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
