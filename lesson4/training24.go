
package main

import (
	"fmt"
	"strings"
)

func anagram(stA, stB string) bool {
	/*
		检查两个字符串是否是anagram
	 */
	num := 0
	for _, char := range stA {
		if strings.ContainsAny(stB, string(char)) {
			num ++
		}
	}

	if num == len(stA) && len(stA) == len(stB) {
		return true
	} else {
		return false
	}
}

func main() {

	var stA, stB string

	fmt.Printf("Enter the first string: ")
	fmt.Scanf("%s", &stA)
	fmt.Printf("Enter the second string: ")
	fmt.Scanf("%s", &stB)

	isAnagram := anagram(stA, stB)

	if isAnagram {
		fmt.Printf("%s and %s are anagrams!", stA, stB)
	} else {
		fmt.Printf("%s and %s are not anagrams!", stA, stB)
	}


}

