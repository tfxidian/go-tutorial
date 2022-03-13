package main

import "fmt"

/**
 * @Description
	new returns a pointer to the struct.
	Each individual field inside the struct can then be accessed
	and its value can be initialized or manipulated.
 * @Author ftang
 * @Date 2022/3/13 15:08
 **/
type Person struct {
	name string
	age int
	gender string
}

type mySlice []int

func main() {
	var person1  = new(Person)
	person1.name = "ftang"
	person1.age = 20
	person1.gender = "male"
	fmt.Println(person1)

	var slice1 = new(mySlice)
	tempslice := append(*slice1, 1,2,3)
	slice1 = &tempslice
	fmt.Println(slice1)
}
