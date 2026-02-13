package main

import "fmt"

func main() {
	// slice定义：
	var s1 []int     // 长度=0，容量=0，零值，nil
	var s2 = []int{} // 长度=0，容量=0，零值，字面量定义
	var s3 = []int{1, 2, 3, 5: 5}
	s4 := make([]int, 5) //	格式：make(类型，size[size=容量])
	s5 := make([]int, 5, 10)
	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s1), cap(s1), s1 == nil)
	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s2), cap(s2), s2 == nil)
	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s3), cap(s3), s3 == nil)
	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s4), cap(s4), s4 == nil)
	fmt.Printf("长度%v: 容量:%v nil:%v \n", len(s5), cap(s5), s5 == nil)
}
