package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p Person
	fmt.Println(p)

	p1 := Person{"yun", "ma", 57}
	fmt.Println(p1)

	p2 := Person{
		FirstName: "jun",
		LastName:  "lei",
		Age:       51,
	}
	fmt.Println("person2", p2)
	p3 := Person{FirstName: "robin"}
	fmt.Println(p3)

	//访问field
	fmt.Println(p1.Age)

	//结构体指针
	ps := &p2
	fmt.Println(*ps)
	fmt.Println((*ps).Age)

	p4 := new(Person)
	p4.LastName = "wang"
	p4.FirstName = "jianlin"
	p4.Age = 63
	fmt.Println(*p4)

	//复制结构体对象 值拷贝
	p5 := p2
	fmt.Println(p5, p2)
	p5.Age = 100
	fmt.Println(p5, p2)

	// Two structs are equal if all their corresponding fields are equal.

}
