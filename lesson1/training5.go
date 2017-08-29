
/*
你将经常编写处理数值的程序。
取决于所使用的编程语言，有时必须将获取的输入转换成数值数据类型。
编写一个程序，提示用户输入两个数。打印这两个数的和、差、积及商，如示例输出所示。
示例输出
What is the first number? 10
What is the second number? 5
10 + 5 = 15
10 - 5 = 5
10 * 5 = 50
10 / 5 = 2

约束
用户输入的值将会是字符串。确保在做数学运算之前将其转换成数值。
将输入和输出与数值转换及其他处理分开。
生成一条在适当的位置包含换行的输出语句。
 */

package main

import (
	"fmt"
)

func main() {
	var firstNum, secondNum int

	fmt.Print("What is the first number? ")
	fmt.Scanf("%d", &firstNum)

	fmt.Print("What is the second number? ")
	fmt.Scanf("%d", &secondNum)

	fmt.Printf("%d + %d = %d\n", firstNum, secondNum, firstNum+secondNum)
	fmt.Printf("%d - %d = %d\n", firstNum, secondNum, firstNum-secondNum)
	fmt.Printf("%d * %d = %d\n", firstNum, secondNum, firstNum*secondNum)
	fmt.Printf("%d / %d = %d\n", firstNum, secondNum, firstNum/secondNum)

}
