package main

import "fmt"

func lowertoupper(alfa rune) rune {
	if alfa >= 'a' && alfa <= 'z' {
		return alfa - ('a' - 'A')
	}
	return alfa
}

func main() {
	var n rune

	fmt.Scanf("%c", &n)

	fmt.Printf("%c", lowertoupper(n))
}
