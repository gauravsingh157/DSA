// package main

// import "fmt"

// func reverseArrays(arr []int) {

// 	start := 0
// 	end := len(arr) - 1

// 	for start < end {
// 		arr[start], arr[end] = arr[end], arr[start]
// 		start++
// 		end--
// 	}

// }

// func main() {

// 	fmt.Println("Before Reverse:")

// }

package main

import "fmt"

func reverseArray(arr []int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		// swap
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println("Before Reverse:", arr)
	reverseArray(arr)
	fmt.Println("After Reverse:", arr)
}
