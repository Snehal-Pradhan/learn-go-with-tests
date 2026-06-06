package main



func Sum(arr []int) int {
	sum := 0
	for _,i:= range arr {
		sum += i
	}
	return sum
}

func SumAll(numberSum ...[]int) []int{
	length := len(numberSum)
	sums := make([]int,length)

	for i,numArr := range numberSum {
		sums[i] = Sum(numArr)
	}
	return sums
}