package main


func Sum(nums []int) int {
	sum := 0
	for _,i := range nums {
		sum += i
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sumArr []int
	for _,val := range numsToSum {
		sumArr =  append(sumArr, Sum(val))
	}
	return sumArr
}

