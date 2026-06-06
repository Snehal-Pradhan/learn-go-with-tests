package main

func Sum(arr []int) int {
	sum := 0
	for _,i:= range arr {
		sum +=i
	}
	return sum
}