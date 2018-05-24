package main

import "fmt"

func main() {
	var a = []int{2, 34, 2, 56, 23, 56, 7, 87, 5, 10}
	var max = findMax(a, 0, len(a)-1)
	fmt.Println("max val :", max)
	fmt.Println("max num in arr :",findMaxNum(a))
}

//Normal Divide and conquer
func findMax(arr []int, firstIndex int, lastIndex int) int {
	if firstIndex == lastIndex {
		return arr[firstIndex]
	}
	mid := (firstIndex + lastIndex) / 2
	max1 := findMax(arr, firstIndex, mid)
	max2 := findMax(arr, mid+1, lastIndex)
	if max1 > max2 {
		return max1
	}
	return max2
}

//go way of divide and conquer
func findMaxNum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	max1 := findMaxNum(arr[:len(arr)/2])
	max2 := findMaxNum(arr[len(arr)/2 :])
	if max1 > max2 {
		return max1
	}
	return max2
}
