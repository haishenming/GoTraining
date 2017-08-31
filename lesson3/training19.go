
/*
经常有这样的需求：看看某个值是不是在特定的范围内，并根据结果来修改程序的流程。
创建一个程序，用一个人的身高（以英寸为单位）和体重（以磅为单位）计算其体质指数（Body  Mass  Index，BMI）。
该程序应提示用户输入身高和体重。
使用如下公式计算BMI：
bmi = (weight / (height × height)) × 703
如果BMI在18.5和25之间，显示这个人为正常体重。如果在该范围之外，则提示用户偏瘦或超重，建议就医咨询。
示例输出
Your BMI is 19.5.
You are within the ideal weight range.
或
Your BMI is 32.5.
You are overweight. You should see your doctor.
约束
确保程序只接受数值数据。如果数据不合法，则不允许继续。
 */

package main

import (
	"fmt"
)

func getBMI(weight, height float64) float64 {
	F64weight := float64(weight)
	F64height := float64(height)
	BMI := (F64weight / (F64height * F64height)) * 703.0

	return BMI
}

func main() {

	var weight, height float64

	fmt.Printf("Weight: ")
	fmt.Scanf("%f", &weight)

	fmt.Printf("Height: ")
	fmt.Scanf("%f", &height)

	BMI := getBMI(weight, height)

	fmt.Printf("Your BMI is %.1f.\n", BMI)

	if BMI < 18.5 {
		fmt.Printf("You are lessweight. You should see your doctor.")
	} else if BMI > 25 {
		fmt.Printf("You are overweight. You should see your doctor.")
	} else {
		fmt.Printf("You are within the ideal weight range.")
	}

}

