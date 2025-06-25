package main

import (
	"fmt"
)

const (
	_ = iota // c0 == 0
	c1
	c2
	c3
	c4
	c5
	c6
)

func main() {
	KB := 1 << (10 * c1) // 1 << 10 == 1024
	MB := 1 << (10 * c2) // 1 << 20 == 1048576
	GB := 1 << (10 * c3) // 1 << 30 == 1073741824
	TB := 1 << (10 * c4) // 1 << 40 == 1099511627776
	PB := 1 << (10 * c5) // 1 << 50 == 1125899906842624
	EB := 1 << (10 * c6) // 1 << 60 == 1152921504606846976
	fmt.Printf("1 KB = %d bytes\n", KB)
	fmt.Printf("1 MB = %d bytes\n", MB)
	fmt.Printf("1 GB = %d bytes\n", GB)
	fmt.Printf("1 TB = %d bytes\n", TB)
	fmt.Printf("1 PB = %d bytes\n", PB)
	fmt.Printf("1 EB = %d bytes\n", EB)
}
