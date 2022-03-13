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
}

func main()  {
	slicedemo()
}


