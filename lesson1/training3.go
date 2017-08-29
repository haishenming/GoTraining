
/*
引号常用于表示一个字符串的开始和结束。
但有时我们需要使用转义字符打印引号本身。
创建一个程序，提示用户输入一条引语及其作者。
按照示例输出那样显示引语和作者。
示例输出:
What is the quote? These aren't the droids you're looking for.
Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids you're looking for."
约束
使用一条输出语句来生成该输出，使用适当的字符串转义技术来处理引语。
如果所用编程语言支持字符串插入或替换，在这个练习中不要使用。要使用字符串连接。
 */

package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	//var quote string
	var name string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the quote? ")
	quote, _, _ := reader.ReadLine()

	//fmt.Print("What is the quote? ")
	//fmt.Scanf("%s", quote)

	fmt.Print("Who said it? ")
	fmt.Scanf("%s\n", &name)

	fmt.Printf("%s says, \"%s\"", strings.Title(name), quote)
}