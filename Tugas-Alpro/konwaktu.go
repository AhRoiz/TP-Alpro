package main

import "fmt"

func main() {
	var h, m, s, hasil int

	fmt.Scan(&h, &m, &s)

	h = h * 60 * 60
	m = m * 60
	s = s

	hasil = h + m + s

	fmt.Print("Hasil Konversi = ", hasil, " detik")
}

func waktu(h, m, s int) int {
	var hasil int

	return hasil
}
