package main

import "fmt"

/*
	单分支：
		条件只能是Bool类型（True/False）
		条件表达式类型一致
		condition条件=True在会执行代码块
*/

func main() {
	var a = 6

	// var b = 2
	// 单分支
	// if a > b {
	// if a, b := 5, 2; a > b {
	// 	fmt.Print("True")
	// }

	// 多分支
	// if a > 6 {
	// 	fmt.Print("大于")
	// } else if a < 6 {
	// 	fmt.Print("小于")
	// } else {
	// 	fmt.Print("相等")
	// }

	// 嵌套写法-不推荐
	if a > 6 {
		fmt.Print("大于")
	} else {
		if a < 6 {
			fmt.Print("小于")
		} else {
			fmt.Print("相等")
		}
	}
}
