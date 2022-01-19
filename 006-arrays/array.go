package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	var y [8]string
	fmt.Println(y)
	var z [3]complex128
	fmt.Println(z)

	x[0] = 100
	x[1] = 101
	x[2] = 102
	x[3] = 103
	x[4] = 105
	fmt.Println(x)

	//初始化方法
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println(a)
	b := [5]int{1, 3, 5, 7, 9}
	fmt.Println(b)
	//不需要全部初始化
	c := [5]int{2}
	fmt.Println(c)
	//不指定数组长度
	d := [...]int{3, 5, 7, 9}
	fmt.Println(d)

	//遍历数组

	name := [3]string{"jack ma", "robin", "pony"}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	m := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)
	for i := 0; i < len(m); i++ {
		sum = sum + m[i]
	}
	fmt.Println(sum)

	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sta", "Sun"}
	for index, value := range daysOfWeek {
		fmt.Printf("Day %d : %s \n", index, value)
	}
	for _, weekday := range daysOfWeek {
		fmt.Printf("weekday: %s\n", weekday)
	}

	twodimension := [2][2]int{
		{3, 5},
		{7, 9},
	}
	fmt.Println(twodimension)

}
