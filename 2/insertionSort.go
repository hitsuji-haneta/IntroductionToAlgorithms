package main

func insertionSort(array []int) {
	for i := 2; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for {
			array[j+1] = array[j]
			j = j - 1
			if j < 0 || array[j] < key {
				break
			}
		}
		array[j+1] = key
	}
}

func main() {

}
