package main

import "fmt"

func main() {
	var a = make([]int, 10)
	a[0], a[1] = 1, 1
	for i := 2; i < 10; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Println(a)
}
