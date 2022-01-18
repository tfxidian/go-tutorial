//package awesomeProject
package main

import "fmt"

func main() {
	fmt.Println("hello, go language!")
	// 定义变量
	var myStr string = "hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)

	//多个定义
	var (
		employeeId          int    = 5
		firstName, lastName string = "fly", "me"
	)
	fmt.Println(employeeId, firstName, lastName)
	//短定义
	name := "jack ma"
	age, salary, isFirstRich := 57, 1, true
	fmt.Println(name, age, salary, isFirstRich)

	// 变量类型
	fmt.Printf("name: %T, age: %T, isFirstRich: %T\n", name, age, isFirstRich)

	//默认值
	var (
		first, last string
		num         int
		floatnum    float64
		isconfirmed bool
	)

	fmt.Printf("fist: %s, last: %s, num: %d, floatnum: %f, isconfirmed: %t\n", first, last, num, floatnum, isconfirmed)
	//Println :可以打印出字符串，和变量
	//Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形
	//也就是说，当需要格式化输出信息时一般选择 Printf，其他时候用 Println 就可以了
	//fmt.Println(first, last, num, floatnum, isconfirmed)
	a := 10
	fmt.Println(a)        //right
	fmt.Printf("%d\n", a) //right
	//fmt.Printf(a)　　//error
}
