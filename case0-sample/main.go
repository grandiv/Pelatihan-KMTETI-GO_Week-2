// CASE 0: Menghitung Panjang Sisi Miring Segitiga

package main

import (
	"fmt"
	"math"
)

func CountLength(a, b float32) float32 {
	return float32(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	a := 2.0
	b := 3.0

	c := CountLength(float32(a), float32(b))
	fmt.Println("length of c:", c)
}
