
/*
只有在想快速猜测一下时才会使用单利。
大部分投资都会使用复利公式，它更为精确。
该公式需要把指数引入到程序中来。
编写一个计算复利的程序。程序会让用户输入本金、投资年限、利率及每年计算利息的次数。
P为本金
r为年利率
t为投资年限
n为每年计算利息的次数
A为到期总额
示例输出
What is the principal amount? 1500
What is the rate? 4.3
What is the number of years? 6
What is the number of times the interest is compounded per year? 4
$1500 invested at 4.3% for 6 years compounded 4 times per year is $1938.84.
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var principal, rate, year, times float64

	fmt.Printf("What is the principal amount? ")
	fmt.Scanf("%f", &principal)

	fmt.Printf("What is the rate? ")
	fmt.Scanf("%f", &rate)

	fmt.Printf("Whatis the number of years? ")
	fmt.Scanf("%f", &year)

	fmt.Printf("Waht is the number of times the interest is compounded per year? ")
	fmt.Scanf("%f", &times)

	investment := principal * math.Pow(1 + rate / 100 / times, times * year)
	investment = math.Ceil(investment * 100.0) / 100.0

	fmt.Printf("$%.2f invested at %f%% for %d years compounded %d times per year is $%.2f.",
				principal, rate, int(year), int(times), investment)
}
