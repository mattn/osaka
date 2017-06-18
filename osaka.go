package osaka

import "strings"

func Encode(s string) string {
	for k, v := range tbl {
		s = strings.Replace(s, k, v, -1)
	}
	return s

}
