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

}
