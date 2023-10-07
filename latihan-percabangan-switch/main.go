package main

import "fmt"

func main() {
	var huruf string
	//huruf = "a"
	huruf = "b"
	switch huruf {
	case "a":
		fmt.Printf("huruf %s adalah huruf vokal", huruf)
	case "i":
		fmt.Printf("huruf %s adalah huruf vokal", huruf)
	case "u":
		fmt.Printf("huruf %s adalah huruf vokal", huruf)
	case "e":
		fmt.Printf("huruf %s adalah huruf vokal", huruf)
	case "o":
		fmt.Printf("huruf %s adalah huruf vokal", huruf)
	default:
		fmt.Printf("huruf %s adalah huruf konsonan", huruf)
	}

}
