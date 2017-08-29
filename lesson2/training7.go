
/*
如果你在一个全球化的环境中工作，则必须同时提供公制和英制单位的信息，而且需要知道何时进行转换，以保证结果最为精确。
编写一个计算房间面积的程序。提示用户输入以英尺为单位的房间的长和宽。
然后显示以平方英尺和平方米为单位的面积。
示例输出
What is the length of the room in feet? 15
What is the width of the room in feet? 20
You entered dimensions of 15 feet by 20 feet.
The area is
300 square feet
27.871 square meters

转换公式为：m**2 = f**2 × 0.09290304

约束
让计算与输出分离。
使用一个常量来保存转换因子。
 */

package main

import (
	"fmt"
)

const factor = 0.09290304

func main() {
	var length, width int

	fmt.Printf("What is the length of the room in feet? ")
	fmt.Scanf("%d", &length)
	fmt.Printf("What is the width of the room in feet? ")
	fmt.Scanf("%d", &width)

	feet := length * width
	meters := float32(feet) * factor

	fmt.Printf("The area is \n%d square feet \n%f square meters", feet, meters)



}
