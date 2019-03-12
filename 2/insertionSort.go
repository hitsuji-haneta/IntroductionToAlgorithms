package main

import "fmt"

func insertionSort(array []int) {
	fmt.Println(array)
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for {
			if j >= 0 && array[j] > key {
				array[j+1] = array[j]
				j = j - 1
			} else {
				break
			}
		}
		array[j+1] = key
	}
}

func main() {
	array := []int{3, 4, 10, 6, 2}
	// array := []int{5, 2, 4, 6, 1, 3}
	insertionSort(array)
	fmt.Println(array)
}
