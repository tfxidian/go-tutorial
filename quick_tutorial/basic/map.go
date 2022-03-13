package main

import "fmt"

/**
 * @Description
 * @Author ftang
 * @Date 2022/3/13 13:55
 **/

func mapdemo()  {
	//初始化

	var map1  = map[int]int{
		1:1,
		2:2,
	}
	fmt.Println(map1)
	map1[1] = 0
	fmt.Println(map1)
	map1[3] = 3
	fmt.Println(map1)
	// make 初始化
	m2 := make(map[string]int)
	fmt.Println(m2)
	if m2 ==nil {
		fmt.Println("m2 is nil")
	}
	m2["one"] = 1
	m2["two"] = 2
	fmt.Println(m2)
	fmt.Println(m2["one"])

	var m3 map[string]int
	if m3 ==nil {
		fmt.Println("m3 is nil")
	}
	//m3["a"] =1
	//m3["b"] =2
	//fmt.Println(m3)
}

func main() {
	mapdemo()
}
/*
Assignment to Entry in Nil Map

func main() {
	var agesMap map[string]int

	agesMap["Amanda"] = 25

	fmt.Println(agesMap["Amanda"])
}
The block of code above specifies the kind of map we want (string: int), but doesn’t actually create a map for us to use.
This will cause a panic when we try to assign values to the map
 */


