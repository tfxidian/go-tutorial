package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}
	//初始化
	var m1 = make(map[string]int)
	fmt.Println(m1)
	if m1 == nil {
		fmt.Println("m1 is nil")
	} else {
		fmt.Println("m is not nil")
	}
	m1["one hundred"] = 100
	m1["two hundred"] = 200
	fmt.Println(m1)
	fmt.Println(m1["one hundred"])
	fmt.Println(m1["two hundred"])
	fmt.Println(len(m1))

	var m3 = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}
	//对存在的key修改值，只会覆盖原有的值
	fmt.Println(m3)
	m3["five"] = 50
	fmt.Println(m3)

	//判断是否存在key
	_, ok := m3["six"]
	if ok {
		fmt.Println("six exists in m3")
	} else {
		fmt.Println("six not exists in m3")
	}
	fmt.Println(ok)

	_, flag1 := m3["five"]
	if flag1 {
		fmt.Println(flag1)
	}
	//只用一个返回值表达式时返回的就是key对应的值
	num := m3["five"]
	fmt.Println(num)

	//删除key
	delete(m3, "five")
	fmt.Println(m3)
	//map是引用类型，当将一个map赋给另一个变量时，它们都引用相同的底层数据结构，
	//因此，一个变量所做的改变将对另一个变量可见
	var m4 = m3
	fmt.Println(m4)
	m3["one"] = 10
	fmt.Println(m3)
	fmt.Println(m4)

	//遍历
	for numstr, num := range m3 {
		fmt.Println(numstr, num)
	}
}
