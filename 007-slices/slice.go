package main

import "fmt"

func main() {
	// slice和array的不同是前面是[], array前面是[len]或[...]
	var s1 = []int{3, 5, 7, 9, 11, 13, 17}
	t := []int{2, 4, 6, 8, 10}
	fmt.Println(s1)
	fmt.Println(t)

	//利用array创建slice, 左闭右开
	var a = [5]string{"alpha", "beta", "gamma", "delta", "epsilon"}
	var slice []string = a[1:4]
	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", slice)
	slice1 := a[1:4]
	slice2 := a[:3]
	slice3 := a[2:]
	slice4 := a[:]
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Println("slice4:", slice4)
	//利用slice 创建slice
	cities := []string{"New York", "Chicago", "Beijing", "Shanghai", "Chengdu"}
	ChinaCities := cities[2:]
	UsCities := cities[0:2]
	fmt.Println(ChinaCities)
	fmt.Println(UsCities)
	cities[3] = "Xian"
	fmt.Println(cities)
	//更改了cities， ChinaCities内容也改变了
	fmt.Println(ChinaCities)
	fmt.Printf("UsCities = %v, len = %d, cap = %d\n", UsCities, len(UsCities), cap(UsCities))
	fmt.Printf("ChinaCities = %v, len = %d, cap = %d\n", ChinaCities, len(ChinaCities), cap(ChinaCities))

	// len 和 cap 对比
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original Slice")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[1:5]
	fmt.Println("\nAfter slicing from index 1 to 5")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[:8]
	fmt.Println("\nAfter extending the length")
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	s = s[2:]
	fmt.Println("\nAfter dropping the first two elements") //slice 复制
	src := []string{"sublime", "vscode", "notepad", "vim"}
	fmt.Println(src)
	//返回的是一个长度为2的slice
	dest := make([]string, 2)
	fmt.Println(dest)
	numElementCopied := copy(dest, src)
	fmt.Println(dest)
	fmt.Println(numElementCopied)

	src2 := append(src, "neovim", "eclipse", "goland")
	fmt.Println(src2)
	fmt.Println(len(src2), cap(src2))
	src[0] = "vi"
	fmt.Println(src)
	fmt.Println(src2)

	// Appending to a slice that has enough capacity to accommodate new elements
	slice7 := make([]string, 3, 10)
	copy(slice7, []string{"C", "C++", "Java"})

	slice8 := append(slice7, "Python", "Ruby", "Go")

	fmt.Printf("slice7 = %v, len = %d, cap = %d\n", slice7, len(slice7), cap(slice7))
	fmt.Printf("slice8 = %v, len = %d, cap = %d\n", slice8, len(slice8), cap(slice8))

	slice7[0] = "C#"
	fmt.Println("\nslice7 = ", slice7)
	fmt.Println("slice8 = ", slice8)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	//0 值 slice
	var s2 []int
	fmt.Printf("s = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

}
