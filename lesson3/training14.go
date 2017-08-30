
/*
并不总是需要复杂的控制架构来解决简单的问题。
有时候，某个程序只是在一种情况下需要多走一步，而其他情况下什么都不用做。
编写一个简单的程序，计算订单税额。
该程序应该提示用户输入订单金额和所在州。
如果是威斯康星州（输入“WI”），
必须加收5.5%的税。
对于威斯康星州的居民，该程序应该显示订单金额、税额和总金额；对于其他居民，只显示总金额。
示例输出
What is the order amount? 10
What is the state? WI
The subtotal is $10.00.
The tax is $0.55. The total is $10.55.
或
What is the order amount? 10
What is the state? MN
The total is $10.00
约束
仅使用一个简单的if语句实现该程序，不使用else子句。
确保美分部分向上取整。
在程序的最后，使用单条输出语句来显示程序的结果。
 */

package main

import (
	"fmt"
	"math"
)

func main() {

	var amount float64
	var state string

	fmt.Printf("What is the order amount? ")
	fmt.Scanf("%f", &amount)

	fmt.Printf("What is the state? ")
	fmt.Scanf("%s", &state)

	fmt.Printf("The subtotal is $%.2f.", amount)

	// 美分取整
	addAmount := math.Ceil(amount*0.055*100.0) / 100.0

	if state == "WI" {
		fmt.Printf("The tax is $%.2f. The total is $%.2f.",
			addAmount, amount+addAmount)
	}

}

