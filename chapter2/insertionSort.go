package main

import "fmt"

func insertionSort(array []int) {
	for index, value := range array {
		j := index - 1
		for {
			if j >= 0 && array[j] > value {
				array[j+1] = array[j]
				j = j - 1
			} else {
				array[j+1] = value
				break
			}
		}
	}
}

func main() {
	// array := []int{3, 4, 10, 6, 2}
	array := []int{5, 2, 4, 6, 1, 3}
	fmt.Println(array)
	insertionSort(array)
	fmt.Println(array)
}
