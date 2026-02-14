package main

import "fmt"

func main() {
	/*
		数组：
			1、长度不可变（固定）
			2、有索引 [k,v] 0:1
			3、元素可变

		注意事项：
			1、不能数字开头
	*/

	// 定义
	var a [3]int                                          // [3]int是类型
	var b = [3]int{}                                      // [3]int{} 零值初始化3个元素的数组
	fmt.Println("a", a, "容器长度:", len(a), "容器容量:", cap(a)) // len代表容器的长度/当前元素的个数
	fmt.Println("b", b, "容器长度:", len(b), "容器容量:", cap(b)) // cap代表容器的容器/底层数组最多能容纳多少个元素

	const n = 3 // 通过常量定义数组
	var c = [n]int{}
	fmt.Println(c)

	var d = [...]int{0, 1, 2, 3, 4, 5} // 通过编译器确认当前大小
	fmt.Println("d:", d, len(d), cap(d))

	fmt.Println("2d array list")
	var dd = [2][3]int{
		{100, 200, 300},
		{1000, 2000, 3000},
	} // [3]int 这是类型 [2]代表是2d数组
	fmt.Println(dd)

	fmt.Println("指定索引初始化")
	var e = []int{
		5: 0,
		4: 1,
		3: 2,
		2: 3,
		1: 4,
		0: 5,
	}
	fmt.Println(
		e[0],
		e[1],
		e[2],
		e[3],
		e[4],
		e[5],
	)

	fmt.Println("索引的修改")
	var f = [...]int{1, 2, 3}
	fmt.Println(f)
	f[0] += 5
	fmt.Println(f)

	println("索引的遍历: 一维")

	// forr
	var g = [...]string{"张三", "李四", "王五"}
	for i, v := range g {
		println(i, v)
	}
	// for
	for i := 0; i < len(g); i++ {
		println(g[i])
	}

}
