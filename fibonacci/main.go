package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++{
		fmt.Print(fibonacci(i)," ")
	}
}

func fibonacci(pos int) int {
	if pos == 0 || pos == 1 {
		return 1
	}
	return fibonacci(pos-1) + fibonacci(pos-2)
}
