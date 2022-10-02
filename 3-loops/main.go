package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, threshold := float64(x), float64(0.1)
	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		
		if guess := z*z; math.Abs(guess-x) < threshold {
			fmt.Printf("Fell within threshold of %v within %v iterations \n", threshold, i)
			break
		}
		
		fmt.Println(z)
	}
	return z
}

func main() {
	count := float64(1)
	fmt.Printf("Square root of %v: %v \n", count, Sqrt(count))
	fmt.Printf("Using builtin function: %v", math.Sqrt(count))
}
