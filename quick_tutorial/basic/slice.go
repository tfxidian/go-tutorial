package main

import "fmt"

func slicedemo(){
	//创建slice
	slice1 := []int{2,4,5,67,9}
	fmt.Println(slice1)

	var slice2 = []int{2,3,6,9,0}
	fmt.Println(slice2)

	//利用array创建slice
	var array1 = [4]int{1,2,3,4}
	var slice3 = array1[1:2]
	fmt.Println(slice3)

	//复制
	src := []int{1,2,3}
	dst := make([]int, 2)
	num := copy(dst, src)
	fmt.Println(dst)
	fmt.Println(num)

	slice4 := make([]string, 4, 10)
	copy(slice4, []string{"a", "b","c","d"})
	fmt.Println(slice4)
	slice5 := append(slice4,"e","f","g")
	fmt.Println(slice4)
	fmt.Println(slice5)

	slice6 := []string{"jack", "john", "peter"}
	slice7 := []string{"bill", "mark", "steve"}
	slice8 := append(slice6, slice7...) //这里有点不一样，不能直接使用slice6
	fmt.Println(slice8)
}

func main()  {
	slicedemo()
}


