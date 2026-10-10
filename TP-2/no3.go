package main

import "fmt"

func cekbilanganM(n, m int) string {
	for n > 0 {
		cek := n % 10
		if cek == m {
			return "YA"
		}
		n /= 10
	}
	return "TIDAK"
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Printf("%s", cekbilanganM(n, m))
}
