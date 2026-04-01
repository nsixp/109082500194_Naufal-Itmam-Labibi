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
