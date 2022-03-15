package main

import "fmt"

/**
 * @Description
 * @Author ftang
 * @Date 2022/3/13 19:45
 **/
type Person2 struct {
	name string
	age int
	gender string
}

func main() {
	p := new(Person2)
	p.name ="a"
	p.age =1
	p.gender = "female"
	fmt.Println(p)
    //var p1 Person2 = new(Person2)
    //上面的写法是错误的，因为左边定义是结构体对象，右边是结构体指针
	var p1 *Person2 = new(Person2)
	p1.name = "b"
	p1.age = 2
	p1.gender = "female"
	fmt.Println(p1)
}