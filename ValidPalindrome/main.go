package main

import (
	"fmt"
	"strings"
)

func main() {

}

func isPalindrome(s string) bool {
	filterStr := ""
	for _, v := range s {
		// fmt.Println(v)x
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57) {
			filterStr += strings.ToLower(string(v))
		}
	}
	fmt.Println(filterStr)
	length := len(filterStr)
	for i, v := range filterStr {
		if v != rune(filterStr[length-1-i]) {
			return false
		}
	}
	return true
}
