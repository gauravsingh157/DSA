package main

import "fmt"

func findMinMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	fmt.Println("Server Started ...")

	arr := []int{5, 2, 9, 1, 7}
	min, max := findMinMax(arr)

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
