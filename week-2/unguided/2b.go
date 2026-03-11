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
