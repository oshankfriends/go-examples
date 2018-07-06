package main

import (
	"fmt"
)

func main() {
	var a = -6
	var shift = 1 << 31
	for i := 0; i<32; i++ {
		if a & shift > 0 {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
		}
		shift >>= 1
	}
}

