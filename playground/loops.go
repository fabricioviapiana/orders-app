package main

import (
	"fmt"
  "math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	last := 1.
	fmt.Println(math.Abs(z-last))
  
	for math.Abs(z-last) > 1e-8 {
		last = z
		z -= (z * z - x) / (2 * z)
		fmt.Printf("z is %f while last is %f\n", z, last)
	}
	return z
}

func main() {
  fmt.Println("result", Sqrt(323223))
  fmt.Println("result 2", math.Sqrt(323223))
	
}