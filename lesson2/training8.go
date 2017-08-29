/*
除法并不总是精确，有时候需要写程序来处理余数，而不是用小数。
编写一个平均分配比萨的程序。提示用户输入就餐人数、比萨数，以及每个比萨分几块。
确保平均分配。显示每个人能得到几块。如果有剩余，显示剩下几块。

示例输出
How many people? 8
How many pizzas do you have? 2
8 people with 2 pizzas
Each person gets 2 pieces of pizza.
There are 0 leftover pieces.
 */

package main

import (
	"fmt"
)

var people, pizzas, pieces int

func splitPizzas(people int, pizzas int, pieces int) (int, int) {

	piecesSum := pizzas * pieces
	piecesForPeople := piecesSum / people
	leftover := piecesSum - piecesForPeople * people

	return piecesForPeople, leftover
}

func addS(num int) string {
	if num>1 {
		return "pisces"
	}else {
		return "pisce"
	}
}

func main() {
	fmt.Printf("How many people? ")
	fmt.Scanf("%d", &people)

	fmt.Printf("How many pieces do each pizza cut?")
	fmt.Scanf("%d", &pieces)

	fmt.Printf("How many pizzas do you have? ")
	fmt.Scanf("%d", &pizzas)

	piecesForPeople, leftover := splitPizzas(people, pizzas, pieces)

	fmt.Printf("Each person gets %d %s of pizza.\nThere are %d leftover %s.",
		piecesForPeople, addS(piecesForPeople), leftover, addS(leftover))
}
