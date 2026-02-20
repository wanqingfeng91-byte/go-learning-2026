package main

import "fmt"

// 定义时 x,y是int类型形参、result是返回值
func add(x, y int) int {
	result := x + y
	return result
}

func main() {
	// 调用时
	r := add(1, 1) // 实参
	fmt.Println(r)
	fmt.Printf("Type: %T", add) // 查看函数的类型
}
