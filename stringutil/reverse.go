package stringutil //name of package must be the same as name of dir

import (
	"bytes"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func PrintRandom(s string) string {
	var buffer bytes.Buffer

	for i := 0; i < 10; i++ {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func PrintRandomAndOne(s1, s2 string) string {
	var buffer bytes.Buffer

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			buffer.WriteString(s1)
		} else {
			buffer.WriteString(s2)
		}
	}
	return buffer.String()
}
