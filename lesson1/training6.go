
/*
你的计算机知道今年的年份，这意味着可以将该信息加到你的程序中。
你只需要知道所用编程语言如何提供该信息。
创建一个程序，确定用户还有多少年退休，并计算退休年份。
它应该提示用户输入当前年龄，以及想要退休的年龄，最终按如下例子这样显示输出。
示例输出
What is your current age? 25
At what age would you like to retire? 65
You have 40 years left until you can retire.
It's 2015, so you can retire in 2055.
约束
再次强调，在做任何数学运算之前，务必将输入转换成数值数据。
不要将当前的年份硬编码到程序中。通过所用编程语言从系统时间中获取。
 */

package main

import (
	"time"
	"fmt"
)

func main() {

	var age, retire int

	fmt.Printf("What is your current age? ")
	fmt.Scanf("%d", &age)
	fmt.Printf("At what age would you like to retire? ")
	fmt.Scanf("%d", &retire)

	thisYear := time.Now().Year()

	fmt.Printf("It's %d, so you can retire in %d.", thisYear, retire-age+thisYear)
}
