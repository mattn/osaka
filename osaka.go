package osaka

import "strings"

func replace(src, s, r string) string {
	dest := ""
	i := 0
	i0 := 0
	for i != -1 {
		i = strings.Index(src[i0:], s)
		if i != -1 {
			dest = dest + src[i0:i] + r
			i += len(s)
			i0 = i
		} else {
			dest = dest + src[i0:len(src)]
		}
	}
	return dest
}

func Encode(s string) string {
	for k, v := range tbl {
		s = replace(s, k, v)
	}
	return s

}
