package iteration

func Repeat(val string) string {
	var res string
	for range 5 {
		res += val
	}
	return res
}
