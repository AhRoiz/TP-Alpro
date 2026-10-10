package main

import "fmt"

func main() {
	var c1, c2, step, celsius, reamur, fahrenheit, kelvin, c, r, f, k, stepC, stepR, stepF, stepK float64
	fmt.Scan(&c1, &c2, &step)

	c = step
	r = 0.8 * step
	f = 1.8*step + 32
	k = step + 273

	stepC = c
	stepR = r
	stepF = f - 32
	stepK = k - 273

	celsius = c1
	reamur = 0.8 * c1
	fahrenheit = 1.8*c1 + 32
	kelvin = c1 + 273

	for celsius <= c2 {
		fmt.Printf("%.2f  |  %.2f  |  %.2f  |  %.2f\n", celsius, reamur, fahrenheit, kelvin)
		celsius += stepC
		reamur += stepR
		fahrenheit += stepF
		kelvin += stepK
	}
}
