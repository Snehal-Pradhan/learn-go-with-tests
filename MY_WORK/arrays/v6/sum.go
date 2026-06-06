package main

func Sum(arr []int) int {
	sum := 0
	for i:=range arr {
		sum += arr[i]
	}
	return sum
}


func SumAllTails(sumArrays ...[]int) []int{
	var result []int
	for _,val := range sumArrays {
		if len(val) == 0 {
			result = append(result, 0)
		}
		result = append(result, Sum(val[1:]))
	}
	return result
}