package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	zold := 0.0
	iter := 0
	for math.Abs(z-zold) > math.Pow(10, -10) {
		zold = z
		z -= (z*z - x) / (2 * z)
		iter++
	}
	fmt.Println(iter)
	return z
}

func main() {
	x := 10.0
	xsqrt := Sqrt(x)
	fmt.Println(xsqrt)
	fmt.Println(math.Abs(xsqrt - math.Sqrt(x)))
}
