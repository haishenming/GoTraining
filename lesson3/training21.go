
/*
很多程序都会以某种形式将信息显示给最终用户，而程序内部则以不同形式表示。
比如，在屏幕上可能会显示“Blue”这个词，而在背后会以某个数值或内部值来表示该颜色。
这是因为程序可能要支持不同语言的用户，比如西班牙语，显示的文本也会有所不同。
编写一个程序，将从1到12的数转换成相应的月份。
提示用户输入数字，显示对应的英文月份：比如输入1，显示January；输入12，则显示December。
如果输入的值超出该范围，则显示相应的错误信息。
示例输出
Please enter the number of the month: 3
The name of the month is March. 
约束
使用switch/case语句实现该程序。
使用单条输出语句。
 */

package main

import (
	"fmt"
)

func main() {

	var num int
	var month string

	fmt.Printf("Please enter the number of the month: ")
	fmt.Scanf("%d", &num)

	switch num := num; num {
	case 1: month = "一月"
	case 2: month = "二月"
	case 3: month = "三月"
	case 4: month = "四月"
	case 5: month = "五月"
	case 6: month = "六月"
	case 7: month = "七月"
	case 8: month = "八月"
	case 9: month = "九月"
	case 10: month = "十月"
	case 11: month = "十一月"
	case 12: month = "十二月"
	default:
		month = ""
	}

	if month != "" {
		fmt.Printf("The name of the month is %s.", month)
	} else {
		fmt.Printf("没有这个月份")
	}

}