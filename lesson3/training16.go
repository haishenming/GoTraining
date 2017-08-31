
/*
除了测试是否相等，还可能需要看一个数与已知值比较是大还是小，并显示相应的消息。
编写一个程序，让用户输入年龄，将其与法定驾驶年龄16比较。
如果用户是16岁，或者更大，程序应该显示“You are old enough to legally drive.”。
如果用户不到16岁，则显示“You are not old enough to legally drive.”。
示例输出
What is your age? 15
You are not old enough to legally drive.
或
What is your age? 35
You are old enough to legally drive.
约束
使用单条输出语句。
使用三元运算符编写该程序。如果所用语言不支持三元运算符，则使用普通的if/else语句，同时仍然使用单条输出语句来显示消息。
 */

package main

import (
	"fmt"
)

func main() {
	const rightfulNumber  = 16  // 合法驾车年龄
	var age int
	fmt.Printf("What is your age? ")
	fmt.Scanf("%d", &age)

	if age >= rightfulNumber {
		fmt.Println("You are old enough to legally drive.")
	} else {
		fmt.Println("You are not old enough to legally drive.")
	}
}
