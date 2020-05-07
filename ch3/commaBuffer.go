package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("12345678910"))
	fmt.Println(commaEnhanced("-1234567"))
	fmt.Println(commaEnhanced("-12345678910.67"))
	fmt.Println(anagram("banana", "nanabana"))
}

func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if !strings.Contains(s2, string(s1[i])) {
			return false
		}
	}
	return true
}

func comma(s string) string {
	var buf bytes.Buffer
	decim := false
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			buf.WriteString(".")
			decim = true
			continue
		}
		if i%3 == 0 && decim == false {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()[1:]
}

func commaEnhanced(s string) string {
	var buf bytes.Buffer
	if s[0] == '-' {
		buf.WriteByte(s[0])
		restOfS := comma(s[1:])
		buf.WriteString(restOfS)
	} else {
		buf.WriteString(comma(s))
	}
	return buf.String()
}
