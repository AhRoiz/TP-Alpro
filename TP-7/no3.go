package main

import "fmt"

const NMAX = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func baca_dan_cetak_array_2() {
	var data tabInt

	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	cetakData(data)
}

func bacaData(A *tabInt) {
	if A.n < NMAX {
		fmt.Scan(&A.info[A.n])
		A.n++
	}
}

func cetakData(A tabInt) {
	for i := 0; i < A.n; i++ {
		fmt.Print(A.info[i], " ")
	}
	fmt.Println()
}

func main() {
	var data tabInt
	for i := 0; i < 6; i++ {
		bacaData(&data)
	}
	cetakData(data)
}
