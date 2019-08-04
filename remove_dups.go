package main

import "fmt"

func main() {
	s := "aabaa"
	fmt.Println(removeDups(s))
}

/** remove duplicate
	any 2 continious same char
	examples
	"aabaa" -> b
	bbxxtt -> ""
	bntbnt -> bntbnt
	"cbe" -> cbe
**/
func removeDups(s string) string {
	var i, j int
	ss := make([]byte, len(s))

	for i = 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] {
			i++
		} else {
			ss[j] = s[i]
			j++
		}
	}
	if s[i] != s[i+1] {
		ss[j] = s[i]
		ss[j+1] = s[i+1]
	}
	return string(ss)
}
