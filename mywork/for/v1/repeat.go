package iteration

func Repeat(a string) string {
	repeated :=""
	for i:=0 ;i < 5;i++ {
		repeated = repeated + a
	}
	return repeated
}