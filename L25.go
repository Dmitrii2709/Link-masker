package main

import (
	"fmt"
	"os"
)

func main() {

	message := os.Args[1]
	fmt.Println(maskingSpam(message))
}

func maskingSpam(a string) string {

	s1 := "http://"
	s2 := " "
	s3 := "*"
	var x []byte

	for i := 0; i < len(a); i++ {
		if len(a) > len(s1) && i < (len(a)-6) && a[i:i+len(s1)] == s1 {
			x = append(x, a[i:i+len(s1)]...)

			for j := i + len(s1); j < len(a); j++ {
				if a[j] != s2[0] {
					x = append(x, s3[0])
				} else {
					x = append(x, s2[0])
					i += len(x) - len(a[:i+1])
					break
				}
			}
			i += len(x) - len(a[:i+1])
		} else {
			x = append(x, a[i])
		}
	}
	return string(x)
}
