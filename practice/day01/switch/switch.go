package main

import (
	"fmt"
)

func main() {
	/*
		Go:	switch与其它语言switch不一样，不会是实现穿透
			想实现穿透的功能
	*/
	var a int = 3
	switch a {
	case 1:
		fmt.Print("1\n")
		fallthrough
	case 2:
		fmt.Print("2\n")
		fallthrough
	case 3:
		fmt.Print("3\n")
		fallthrough
	case 4:
		fmt.Print("4\n")
		fallthrough
	case 5:
		fmt.Print("5\n")
		fallthrough
	default:
		fmt.Print("穿透")
	}

}
