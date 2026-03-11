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
