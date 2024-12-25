package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scanln(&n)

	pi := 0.0
	prevTerm := 0.0
	iteration := 0

	for i := 0; i < n; i++ {
		iteration++
		term := (1.0 - 2.0*float64(i%2)) / float64(2*i+1)
		pi += term

		if math.Abs(term-prevTerm) < 0.00001 {
			fmt.Printf("Hasil PI: %.10f\n", pi*4)
			fmt.Printf("Pada i ke: %d\n", iteration)
			return
		}
		prevTerm = term
	}

	pi *= 4

	fmt.Printf("Hasil PI: %.7f\n", pi)
}
