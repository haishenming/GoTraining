
/*
有时必须向上取整到下一个数，而不是遵循标准的舍入规则。
计算粉刷天花板需要多少加仑涂料。
提示输入长和宽，假设1加仑可以粉刷350平方英尺。
以整数形式显示粉刷该房间的天花板所需的加仑数。
示例输出
You will need to purchase 2 gallons of paint to cover 360 square feet.
记住，涂料必须整加仑购买。所以要向上取整到下一个整数。

约束
使用一个常量来保存转换率。
确保向上取整到下一个整数。
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var lenght, width float64

	fmt.Printf("Lenght: ")
	fmt.Scanf("%f", &lenght)

	fmt.Printf("Width: ")
	fmt.Scanf("%f", &width)

	fmt.Printf("You will need to purchase %d gallons of paint to cover %d square feet.",
				int(math.Ceil(lenght*width/350.0)), int(lenght*width))
}