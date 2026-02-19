package main

import "fmt"

func main() {
	// var c = '测'
	// fmt.Printf("%T,%[1]d,%#[1]x,%[1]c,%[1]q\n", c)

	s := "字符串"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
