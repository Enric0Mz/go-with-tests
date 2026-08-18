package iteration

import "strings"

func Repeat(val string) string {
	var res strings.Builder
	for range 5 {
		res.WriteString(val)
	}
	return res.String()
}
