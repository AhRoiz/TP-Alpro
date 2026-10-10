package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func JumlahdanRataRata() {
	var data tabInt
	var nData int

	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData))
	fmt.Printf("%.1f\n", rataRata(data, nData))
}

func baca(A *tabInt, n *int) {
	*n = 0
	var temp int
	for *n < NMAX {
		fmt.Scan(&temp)
		if temp == 0 {
			break
		}
		if temp < 0 {
			temp = -temp
		}
		(*A)[*n] = temp
		*n++
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += A[i]
	}
	return total
}

func rataRata(A tabInt, n int) float64 {
	return float64(jumlah(A, n)) / float64(n)
}

func main() {
	JumlahdanRataRata()
}
