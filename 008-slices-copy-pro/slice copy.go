package main

import "fmt"

func main() {
	//基本的复制
	src := []int{1, 2, 3}
	dst := make([]int, 2)
	num := copy(dst, src)
	fmt.Println(src)
	fmt.Println(dst)
	fmt.Println(num)

	//append 用法
	slice1 := []int{1, 2, 3, 4}
	fmt.Printf("cap of slice1 is :%d ", cap(slice1))
	slice2 := append(slice1, 5, 6, 7)
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = 0
	fmt.Println(slice1)
	fmt.Println(slice2)
	/*
		append追加元素，如果slice还有容量的话，就会将新的元素放在原来slice后面的剩余空间里，
		当底层数组装不下的时候，Go就会创建新的底层数组来保存这个切片，slice地址也随之改变。
	*/

	slice3 := make([]string, 4, 10)
	copy(slice3, []string{"a", "b", "c", "d"})
	slice4 := append(slice3, "e", "f", "g")
	fmt.Println(slice3)
	fmt.Println(slice4)
	slice3[0] = "x"
	fmt.Println(slice3)
	fmt.Println(slice4)

	//slice 复制
	slice5 := []string{"jack", "john", "peter"}
	slice6 := []string{"bill", "mark", "steve"}
	slice7 := append(slice5, slice6...) //这里有点不一样，不能直接使用slice6

	fmt.Println(slice5)
	fmt.Println(slice6)
	fmt.Println(slice7)

	for _, name := range slice7 {
		fmt.Println(name)
	}
}
