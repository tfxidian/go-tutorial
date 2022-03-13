package main

import "fmt"

/**
 * @Description
 * @Author ftang
 * @Date 2022/3/13 14:34
 **/
func main() {
	var q *int
	fmt.Println("q = ", q)
	var a = 20
	q = &a
	fmt.Println(q)
	fmt.Println(*q)

	p := new(int)
	*p = 100
	fmt.Println(*p)

}

/*

The introduction documents dedicate many paragraphs to explaining the difference between new() and make(), but in practice, you can create objects within local scope and return them.

Why would you use the pair of allocators?
Things you can do with make that you can't do any other way:

Create a channel
Create a map with space preallocated
Create a slice with space preallocated or with len != cap



 */