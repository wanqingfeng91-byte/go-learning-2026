package main

import (
	"fmt"
)

func main() {
	// slice定义：
	/* 	var s1 []int     // 长度=0，容量=0，零值，nil
	   	var s2 = []int{} // 长度=0，容量=0，零值，字面量定义
	   	var s3 = []int{1, 2, 3, 5: 5}
	   	s4 := make([]int, 5) //	格式：make(类型，size[size=容量])
	   	s5 := make([]int, 5, 10)
	   	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s1), cap(s1), s1 == nil)
	   	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s2), cap(s2), s2 == nil)
	   	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s3), cap(s3), s3 == nil)
	   	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s4), cap(s4), s4 == nil)
	   	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s5), cap(s5), s5 == nil) */

	/*
		slice原理：
			元数据 = strust 结构体
			type slice type strust {
				array unsafe.Pointer
				len int
				cap int
				1、slice 本身就是一个结构体值（不是指针）
				2、结构体中的 array 指针指向底层数组
				3、多个 slice 可以指向同一个底层数组
			}
	*/
	s1 := []int{100}                                                                                     // []int类型的切片
	fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %p\n", s1, &s1, len(s1), cap(s1), &s1[0]) // %p 打印的是在内存中的（指针）地址
	s1 = make([]int, 2, 5)                                                                               // []int整数类型的切片、长度2、容量5
	fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %p\n", s1, &s1, len(s1), cap(s1), &s1[0]) // %p 打印的是在内存中的（指针）地址
	// 思考一个问题： 上述31行相当于重新赋值，底层数组地址 或者 切片遍历的地址是否会变化？

}
