package main

import "fmt"

func main() {
	num := 4
	if num%2 == 0 {
		fmt.Printf("%d Merupakan bilangan genap\n", num)
	} else {
		fmt.Printf("%d Merupakan bilangan ganjil\n", num)
	}
	num = 5
	if num%2 == 0 {
		fmt.Printf("%d Merupakan bilangan genap", num)
	} else {
		fmt.Printf("%d Merupakan bilangan ganjil", num)
	}
}
