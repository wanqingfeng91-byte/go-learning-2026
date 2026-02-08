// 这是包注释
package main

import (
	"fmt"
)

// TODO: 将来完成
func main() {
	// fmt.Println("Hello,World")
	// const a int = 100
	// const (
	// 	b = "1" // 不指定类型会自动类型推导
	// 	c = 2
	// 	d = '3'
	// )

	// const a = iota
	// const b = iota
	// fmt.Println(a, b)

	// 批量学iota是从0开始计时
	// const (
	// 	c = iota // 0
	// 	d = iota // 1
	// 	e = iota // 2
	// 	f = iota // 3
	// )

	// fmt.Println(c, d, e, f)

	// const (
	// 	g = iota
	// 	h
	// 	i
	// 	j
	// )

	// fmt.Println(g, h, i, j)

	// var a int = 1
	// fmt.Println(a)

	// var b, c, d, e, f int
	// fmt.Println(b, c, d, e, f)
	// var b = 200
	// fmt.Print(b)

	var a int = 100
	var b = 200
	var c = 300
	var d, e int = 1, 2
	fmt.Println(a, b, c, d, e)
}
