package main

import "fmt"

var hargaA, hargaB, hargaC int16 = 10000, 15000, 7000

func main() {
	totC := hargaA + hargaB + hargaC
	fmt.Printf("Total harga belanja : %d", totC)
}
