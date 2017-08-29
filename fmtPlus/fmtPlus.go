package fmtPlus

import (
	"fmt"
	"bufio"
	"os"
)



func Input(tip string) string {

	fmt.Print(tip)
	reader := bufio.NewReader(os.Stdin)
	ret, _, _ := reader.ReadLine()

	return string(ret)
}

//func RawInput(tip string) {
//
//	fmt.Print(tip)
//	reader := bufio.NewReader(os.Stdin)
//	ret, _, _ := reader.ReadLine()
//
//	return ret
//}
