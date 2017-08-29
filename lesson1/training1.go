
/*
在很多编程语言中，“Hello, World”程序都是你要学习编写的第一个程序，但是它没有涉及任何输入。
所以先来创建一个提示输入名字并打印包含该名字的问候信息的程序。
示例输出:
What is your name? Brian
Hello, Brian, nice to meet you!

约束
输入、字符串连接和输出这几个部分要分开。
 */


package main

import (
	"fmt"
)

func main() {

	var name string

	fmt.Printf("What is your name? ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello, %s, nice to meet you!", name)
}


// 挑战：不使用任何变量，编写一个新版本的程序。