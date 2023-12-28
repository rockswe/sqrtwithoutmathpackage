package main

import (
	"fmt"
)

// Sqrt calculates the square root of x.
func Sqrt(x float64) float64 {
	if x < 0 {
		return -1 // Return -1 for negative numbers
	}
	if x == 0 {
		return 0
	}

	z := x / 2 // Initial guess
	for i := 0; i < 1000; i++ {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		if z == prevZ {
			break // Break if the value stops changing
		}
	}
	return z
}

func main() {
	number := 64.0
	fmt.Printf("The square root of %v is %v\n", number, Sqrt(number))
	fmt.Printf("unknown pass")
}
