package main


func Sum(arr [5]int) int {
	sum := 0

	for i := range arr {
		sum += arr[i]
	}

	return sum

}