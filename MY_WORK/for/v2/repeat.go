package iteration

const repeatCount = 5

func Repeat(a string) string {
	repeated := ""
	for i:= 0;i<repeatCount;i++ {
		repeated += a
	}
	return repeated
}