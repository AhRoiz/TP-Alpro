package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func baca_dan_cetak_array_1() {
	var data tabInt
	var banyakData int

	bacaData(&data, &banyakData)
	cetakData(data, banyakData)
}

func bacaData(A *tabInt, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&(*A)[i])
	}
}

func cetakData(A tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
	} else {
		for i := 0; i < n; i++ {
			fmt.Print(A[i], " ")
		}
		fmt.Println()
	}
}

func main() {
	baca_dan_cetak_array_1()
}
