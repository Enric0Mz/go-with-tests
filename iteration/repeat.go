package iteration

import "strings"

func Repeat(val string, repeat int) string {
	var res strings.Builder
	for range repeat {
		res.WriteString(val)
	}
	return res.String()
}
