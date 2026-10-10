package main

import "fmt"

func main() {
	var n float64
	var hasil float64
	var member bool

	fmt.Scan(&n, &member)

	hasil = n

	if n > 100000 && member == true {
		hasil = hasil - (n * 0.1)
	} else if n > 100000 && member == false {
		hasil = hasil - (n * 0.05)
	}

	fmt.Print(hasil)
}
