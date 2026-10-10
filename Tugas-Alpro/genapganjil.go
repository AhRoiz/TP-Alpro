package main

import "fmt"

func main() {
	var n, genap, ganjil int

	fmt.Print(" Masukkan nilai : ")
	fmt.Scan(&n)

	ganjil = 0
	genap = 0

	for n > 0 {
		digit := n % 10
		if digit/2 == 0 {
			genap++
		} else {
			ganjil++
		}
		n /= 10
	}

	fmt.Print(ganjil, " ", genap)
}
