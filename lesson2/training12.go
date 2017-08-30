
/*
计算单利是快速确定投资是否有价值的不错方式。
它也能让你熟悉如何在程序中显式控制运算顺序。
编写一个计算单利的程序。提示输入本金、利率及时间，显示到期总额（本金+利息）。
单利的公式是A = P(1 + rt)，
其中P为本金，r为年利率，t为投资年限，A是这笔投资到期后的总额。
示例输出
Enter the principal: 1500
Enter the rate of interest: 4.3
Enter the number of years: 4
After 4 years at 4.3%, the investment will be worth $1758.
约束
提示利率输入的是百分比（像15，而不是0.15）。在程序中要除以100。
对于小数部分，不足1美分的要向上取整。确保输出格式化为货币形式。
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var principal, rate, year float64

	fmt.Printf("Enter the principal: ")
	fmt.Scanf("%f", &principal)
	fmt.Printf("Enter the reta of interest: ")
	fmt.Scanf("%f", &rate)
	fmt.Printf("Enter the number of years: ")
	fmt.Scanf("%f", &year)


	investment := principal * (1 + rate / 100 * year)
	investment = math.Ceil(investment * 100.0) / 100.0


	fmt.Printf("After %d years at %.2f%%, the investment will be worth $%.2f.", int(year), rate, investment)


}
