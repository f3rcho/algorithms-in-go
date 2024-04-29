package main

import "fmt"

func main() {
	arr := []int32{5, 2, 4, 6, 1, 3}
	result := CountingSort(arr)
	fmt.Println(result)
}

func CountingSort(arr []int32) []int32 {
	countingArr := make([]int32, 100)

	for _, num := range arr {
		countingArr[num]++
	}
	return countingArr
}
