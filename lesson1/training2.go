
/*
创建一个程序，提示用户输入字符串，然后输出这个字符串，以及其中包含的字符数。
示例输出:
What is the input string?
Homer Homer has 5 characters.

约束
确保输出中包含原始的字符串。
使用一个输出语句来构造输出。
使用所用编程语言中的一个内置函数来确定字符串的长度。
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	for {
		fmt.Printf("What is the input string? ")
		fmt.Scanf("%s", &str)
		if str != "" {
			fmt.Printf("%s has %d characters.", strings.Title(str), len(str))
			break
		}else {
			fmt.Println("Please input a string!")
		}
	}



}
