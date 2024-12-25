package main

import (
	"fmt"
)

func main() {
	var sum, count float64
	var bilangan float64

	fmt.Println("Masukkan bilangan riil (akhiri dengan 9999):")

	for {
		fmt.Scan(&bilangan)
		if bilangan == 9999 {
			break
		}
		sum += bilangan
		count++
	}

	if count == 0 {
		fmt.Println("Tidak ada bilangan yang dimasukkan.")
	} else {
		average := sum / count
		fmt.Printf("Rerata dari bilangan yang dimasukkan adalah: %.2f\n", average)
	}
}
