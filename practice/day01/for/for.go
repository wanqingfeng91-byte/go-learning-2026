package main

// "time"

func main() {
	// for 循环/*  */
	// for i := 100; ; /* i < 105 */ /* i++ */ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	/*
		第一次：i=100；对比循环条件;结果为True，执行代码块;则i++
		第二次：i=101；对比循环条件;结果为True，执行代码块;则i++
		第三次：i=102；对比循环条件;结果为True，执行代码块;则i++
		第四次：i=103；对比循环条件;结果为True，执行代码块;则i++
		第五次：i=104；对比循环条件;结果为True，执行代码块;则i++
		第六次：i=105；对比循环条件;结果为False，循环中断

		注释循环条件 OR i++ 就是死循环
	*/

	// continue：跳过本次循环
	/* 	for i := 0; i <= 5; i++ {
		if i == 4 {
			continue
		}
		fmt.Println(i)
	} */

	// break：打断循环
	/* 	for i := 0; i <= 5; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	} */

}
