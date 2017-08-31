
/*
我们经常需要确定程序的哪一部分应该根据用户的输入或其他事件来运行。
创建一个程序，将温度从华氏温度（F）转换成摄氏温度（C），或者相反。
提示用户输入起始温度。
程序应该提示转换类型并执行转换。
转换公式为：
C = (F − 32) × 5 / 9
和
F = (C × 9 / 5) + 32
示例输出
Press C to convert from Fahrenheit to Celsius.
Press F to convert from Celsius to Fahrenheit.
Your choice: C
Please enter the temperature in Fahrenheit: 32
The temperature in Celsius is 0.
约束
确保支持大小写的C和F。
使用的输出语句要尽可能少，避免输出字符串重复。
 */

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func temperatureTran(temperature int, temperatureType string) string {

	var ret string = "不能识别的温度类型"

	temperatureType = strings.ToUpper(temperatureType)

	if temperatureType == "C" {
		temper := (temperature * 9 / 5) + 32
		ret = strconv.Itoa(temper) + "F"
	} else if temperatureType == "F" {
		temper := (temperature - 32) * 5 / 9
		ret = strconv.Itoa(temper) + "C"
	}

	return ret
}

func main() {

	var temperNum int
	var temperType string

	fmt.Printf("Press C to convert from Fahrenheit to Celsius.\nPress F to convert from Celsius to Fahrenheit.\n")
	fmt.Printf("Please enter the temperature in Fahrenheit: ")
	fmt.Scanf("%d", &temperNum)
	fmt.Scanf("%s", &temperType)

	ret := temperatureTran(temperNum, temperType)

	fmt.Println(ret)
}
