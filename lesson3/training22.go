
/*
将输入和一个已知的值比较，非常普通。
更多时候，我们需要处理一组输入。编写一个程序，让用户输入3个数。
首先确认所有数字各不相同。
如果存在相同的数，退出程序，否则显示其中最大的。
示例输出
Enter the first number: 1
Enter the second number: 51
Enter the third number: 2
The largest number is 51.
约束
手动编写该算法。不要使用寻找列表中最大值的内置函数。
 */

package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var biggerNumber, biggestNumber int

	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("Enter the third number: ")
	fmt.Scanf("%d", &c)

	if a > b {
		biggerNumber = a
	} else {
		biggerNumber = b
	}

	if biggerNumber > c {
		biggestNumber = biggerNumber
	} else {
		biggestNumber = c
	}

	fmt.Printf("The largest number is %d", biggestNumber)


}
