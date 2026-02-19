package main

import "fmt"

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
	// s1 := []int{100}                                                                                     // []int类型的切片
	// fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %p\n", s1, &s1, len(s1), cap(s1), &s1[0]) // %p 打印的是在内存中的（指针）地址
	// s1 = make([]int, 2, 5)                                                                               // []int整数类型的切片、长度2、容量5
	// fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %p\n", s1, &s1, len(s1), cap(s1), &s1[0]) // %p 打印的是在内存中的（指针）地址
	// fmt.Println(s1)
	// // 思考一个问题： 上述31行相当于重新赋值，底层数组地址 或者 切片遍历的地址是否会变化？
	// fmt.Println("***********************************")
	// // _ = append(s1, 200)
	// s1 = append(s1, 200)
	// fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s1, &s1, len(s1), cap(s1), &s1) // %p 打印的是在内存中的（指针）地址
	// s2 := append(s1, 1, 2)
	// fmt.Printf("s2: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s2, &s2, len(s2), cap(s2), s2) // %p 打印的是在内存中的（指针）地址
	// // s1 & s2 共用底层数组
	// // 问题：如果修改s2[0]的元素变化 s1和s2 是否一致？
	// s2[0] = 111
	// fmt.Println(s1)
	// fmt.Println(s2)
	// fmt.Println("--------------------------")
	// s3 := append(s1, 1111)
	// fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s1, &s1, len(s1), cap(s1), &s1) // %p 打印的是在内存中的（指针）地址
	// fmt.Printf("s2: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s2, &s2, len(s2), cap(s2), &s2) // %p 打印的是在内存中的（指针）地址
	// fmt.Printf("s3: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s3, &s3, len(s3), cap(s3), &s3) // %p 打印的是在内存中的（指针）地址
	// fmt.Println("--------------------------")
	// s4 := append(s3, 3, 4, 5)
	// fmt.Printf("s1: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s1, &s1, len(s1), cap(s1), &s1) // %p 打印的是在内存中的（指针）地址
	// fmt.Printf("s2: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s2, &s2, len(s2), cap(s2), &s2) // %p 打印的是在内存中的（指针）地址
	// fmt.Printf("s3: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s3, &s3, len(s3), cap(s3), &s3) // %p 打印的是在内存中的（指针）地址
	// fmt.Printf("s4: 底层数组地址:%p, 切片变量地址:%p,长度:%v,容量:%v 底层数组变量地址: %v\n", s4, &s4, len(s4), cap(s4), &s4) // %p 打印的是在内存中的（指针）地址

	// fmt.Println("截取自切片")
	// s1 := []int{1, 3, 5, 7, 9}
	// // 循环取s1切片的元素的地址
	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("%d : addr = %p\n", i, &s1[i])
	// }
	// fmt.Println("------------------")
	// fmt.Printf("s1 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	// s2 := s1 // s2 == s1 完全引用s1
	// fmt.Printf("s2 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	// s3 := s1[:] // s3 == s1 完全引用s1
	// fmt.Printf("s3 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	// fmt.Println("------------------")
	// s4 := s1[1:] // 掐头,长度=4,容量=4,首地址偏移1个元素,底层数组一致s1[1]=s4[0]
	// fmt.Printf("s4 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s4, &s4[0], len(s4), cap(s4), s4)
	// s5 := s1[1:4] // 掐头、去尾3, 5, 7, 9 公用
	// fmt.Printf("s5 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s5, &s5[0], len(s5), cap(s5), s5)
	// s6 := s1[:4] // 去尾 公用
	// fmt.Printf("s6 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s6, &s6[0], len(s6), cap(s6), s6)
	// fmt.Println("------------------")
	// s7 := s1[1:1] // 左包 右不包
	// // fmt.Printf("s7 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s7, &s7[0], len(s7), cap(s7), s7)
	// s7 = append(s7, 111)
	// fmt.Printf("s7 header:%p,数组地址:%p,l=%-2d,c=%-2d,%v\n", &s7, &s7[0], len(s7), cap(s7), s7)
	// fmt.Println("------------------")

	fmt.Println("切片的其它操作")
	s1 := []int{10, 20, 30}
	fmt.Printf("s1 header: %p,数组地址: %p,l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	var s2 []int
	// copy函数
	n := copy(s2, s1)
	fmt.Printf("s2 header: %p,l=%-2d,c=%-2d,%v\n", &s2, len(s2), cap(s2), s2)
	fmt.Print(n, "\n") // 长度为0 没有拷贝
	s2 = append(s2, 1) // 1
	fmt.Printf("s2 header: %p,数组地址: %p,l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Println("------------------")
	n = copy(s2, s1[1:])
	s2[0] += 1
	fmt.Print(n, "\n") // 长度为1
	fmt.Printf("s1 header: %p,数组地址: %p,l=%-2d,c=%-2d,%v\n", &s1, &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("s2 header: %p,数组地址: %p,l=%-2d,c=%-2d,%v\n", &s2, &s2[0], len(s2), cap(s2), s2)
	fmt.Println("------------------")

	// 合并
	s3 := make([]int, len(s1)+len(s2))
	fmt.Printf("s3 header: %p, %p, l=%-2d, c=%-2d, %v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	copy(s3, s1)
	fmt.Printf("s3 header: %p, %p, l=%-2d, c=%-2d, %v\n", &s3, &s3[0], len(s3), cap(s3), s3)
	copy(s3[len(s1):], s2)
	fmt.Printf("s3 header: %p, %p, l=%-2d, c=%-2d, %v\n", &s3, &s3[0], len(s3), cap(s3), s3)

}
