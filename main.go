package main

import (
	"fmt"
	"math"
)

func diagonalDiference(arr [][]float64) float64 {
	var primaryDiagonal, secondaryDiagonal float64
	for i := 0; i < len(arr); i++ {
		primaryDiagonal += arr[i][i]
		fmt.Println(arr[i][i], "pri")
		secondaryDiagonal += arr[i][len(arr)-i-1]
		fmt.Println(arr[i][len(arr)-i-1], "sec")
	}
	return math.Abs(primaryDiagonal - secondaryDiagonal)
}

func main() {
	arr := [][]float64{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	fmt.Println(diagonalDiference(arr)) // Output: 15.0
}

// func main() {
// 	a := []int{2, 4, 6, 4, 2, 7, 6}
// 	uniqueElement := 0

// 	// Loop through the array and find the unique element using XOR
// 	for i := 0; i < len(a); i++ {
// 		uniqueElement ^= a[i]
// 	}

// 	fmt.Println("Unique Element:", uniqueElement) // Output: 7
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func miniMaxSum(arr []int) {
// 	// Sorting the array in ascending order
// 	sort.Ints(arr)

// 	// Calculating the sum of the four smallest numbers
// 	var minSum int
// 	for _, v := range arr[:4] {
// 		minSum += v
// 	}

// 	// Calculating the sum of the four largest numbers
// 	var maxSum int
// 	for _, v := range arr[1:] {
// 		maxSum += v
// 	}

// 	fmt.Println(minSum, maxSum)
// }

// func main() {
// 	arr := []int{1, 3, 5, 7, 9}
// 	miniMaxSum(arr)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func timeConversion(s string) string {
// 	hour, _ := strconv.Atoi(s[:2])
// 	indicator := s[len(s)-2:]

// 	if indicator == "AM" && hour == 12 {
// 		hour = 0
// 	} else if indicator == "PM" && hour != 12 {
// 		hour += 12
// 	}

// 	hourString := fmt.Sprintf("%02d", hour)
// 	convertedTime := hourString + s[2:len(s)-2]
// 	return convertedTime
// }

// func main() {
// 	timeString := "07:45:32PM"
// 	convertedTime := timeConversion(timeString)
// 	fmt.Println(convertedTime) // Output: "19:45:32"
// }
