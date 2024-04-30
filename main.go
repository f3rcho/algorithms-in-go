package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)
	a := x[4:]
	y := alterSlice(a)
	fmt.Println(x)
	fmt.Println(y)
}

func alterSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

// func diagonalDiference(arr [][]float64) float64 {
// 	var primaryDiagonal, secondaryDiagonal float64
// 	for i := 0; i < len(arr); i++ {
// 		primaryDiagonal += arr[i][i]
// 		fmt.Println(arr[i][i], "pri")
// 		secondaryDiagonal += arr[i][len(arr)-i-1]
// 		fmt.Println(arr[i][len(arr)-i-1], "sec")
// 	}
// 	return math.Abs(primaryDiagonal - secondaryDiagonal)
// }

// func main() {
// 	ch := make(chan *int, 4)
// 	array := []int{1, 2, 3, 4}
// 	wg := sync.WaitGroup{}
// 	// Add the number of works equal to the number of array elements.
// 	wg.Add(len(array))
// 	go func() {
// 		for _, value := range array {
// 			ch <- &value
// 		}
// 	}()
// 	go func() {
// 		for value := range ch {
// 			fmt.Println(*value)
// 			// Decrement the waitgroup counter with each iteration.
// 			wg.Done()
// 		}
// 	}()

// 	wg.Wait()
// }

// func main() {
// 	audience := []sorting.Human{
// 		{Name: "Alice", Age: 35},
// 		{Name: "Bob", Age: 45},
// 		{Name: "James", Age: 25},
// 	}

// 	sort.Sort(sorting.AgeFactor(audience))
// 	// fmt.Println(audience)
// 	containStr := strings.Contains("hello world", "wor")
// 	fmt.Println(containStr)

// 	listObj := []int{1, 2, 3}
// 	swap.SwapContents(listObj)
// 	// fmt.Println(listObj)
// 	// arr := [][]float64{
// 	// 	{11, 2, 4},
// 	// 	{4, 5, 6},
// 	// 	{10, 8, -12},
// 	// }
// 	// fmt.Println(diagonalDiference(arr)) // Output: 15.0
// }

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
