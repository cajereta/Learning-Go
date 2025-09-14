package main

import "fmt"

func main() {
	var b byte
	var smallI int32
	var bigI uint64

	b = 255
	smallI = 2147483647
	bigI = 18446744073709551615

	// Addition here
	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("Here is b: %d, smallI: %d and bigI: %d", b, smallI, bigI)

}
