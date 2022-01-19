package main

import "fmt"

func main() {
	//指针基本用法
	var q *int
	fmt.Println("q = ", q)

	var a = 5.75
	var p = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	*p = 200
	fmt.Println("a: ", a)
	fmt.Println("*p: ", *p)
	//new的用法
	ptr := new(int)
	*ptr = 100
	fmt.Println(*ptr)

	var x = 67
	var ptr2 = &x
	//go语言指针是不支持运算的，也就是不支持++/--操作，但是我们借助于unsafe包，可以完成这个操作(这里暂时不涉及)
	//var p1 = p+1
	//ptr2++
	fmt.Println(ptr2)
}
