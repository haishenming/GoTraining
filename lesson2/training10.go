
/*
在处理多个输入和货币时，可能会引入一些微妙的精度问题。
创建一个简单的自助结账系统。提示输入3种商品各自的价格和数量。
计算税前总价。然后以5.5%的税率计算税额。
打印商品的数量和总价，然后打印税前总价、税额和总价。
示例输出
Enter the price of item 1: 25
Enter the quantity of item 1: 2
Enter the price of item 2: 10
Enter the quantity of item 2: 1
Enter the price of item 3: 4
Enter the quantity of item 3: 1
Subtotal: $64.00 Tax: $3.52 Total: $67.52
约束
将程序的输入、处理和输出各部分分开。收集输入，进行数学运算并构建字符串，然后打印输出。
在进行任何计算之前，务必显式地将输入转换为数值数据。
 */

package main

import (
	"fmt"
)

func main() {
	var subtotal, price float64
	var num int

	fmt.Printf("How many items would you want? ")
	fmt.Scanf("%d", &num)

	for i:=0;i<num;i++{
		fmt.Printf("Enter the price of item %d: ", i+1)
		fmt.Scanf("%f", &price)
		subtotal += price
	}

	fmt.Printf("Subtotal: $%.2f Tax: $%.2f Total: $%.2f",
				subtotal, subtotal*5.5/100.0, subtotal+subtotal*5.5/100.0)

}