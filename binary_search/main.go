package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := []int64{}
	var dimension int
	fmt.Println("Please enter array dimension")
	fmt.Scanf("%d", &dimension)
	for i := 0; i < dimension; i++ {
		array = append(array, int64(i+1))
	}
	num := rand.Int63n(int64(dimension - 1))
	fmt.Printf("We are searching for: %d\n", num)
	count, mid := binary_search(array, num)
	fmt.Printf("Number of iterations: %d, we find the number: %d\n", count, mid)
}

func binary_search(array []int64, to_search int64) (count int, mid int64) {
	var low, high int64
	low = 0
	high = int64(len(array)) - 1
	count = 1
	for low <= high {
		mid = (low + high) / 2
		fmt.Printf("low: %d, high: %d, mid: %d, count: %d\n", (low + 1), (high + 1), (mid + 1), count)
		if array[mid] == to_search {
			break
		}
		if array[mid] > to_search {
			high = mid - 1
		} else {
			low = mid + 1
		}
		count++
	}
	mid++
	return
}
