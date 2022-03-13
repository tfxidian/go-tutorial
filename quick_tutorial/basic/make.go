package main

/**
 * @Description the usage of make()
	golang 分配内存主要有内置函数new和make，
	make只能为slice, map, channel分配内存，并返回一个初始化的值。首先来看下make有以下三种不同的用法：
	1. make(map[string]string)
	2. make([]int, 2)
	3. make([]int, 2, 4)
	1. 第一种用法，即缺少长度的参数，只传类型，这种用法只能用在类型为map或chan的场景，例如make([]int)是会报错的。这样返回的空间长度都是默认为0的。
	2. 第二种用法，指定了长度，例如make([]int, 2)返回的是一个长度为2的slice
	3. 第三种用法，第二参数指定的是切片的长度，第三个参数是用来指定预留的空间长度，例如a := make([]int, 2, 4), 这里值得注意的是返回的切片a的总长度是4，预留的意思并不是另外多出来4的长度，其实是包含了前面2个已经切片的个数的。所以举个例子当你这样用的时候 a := make([]int, 4, 2)，就会报语法错误。
 * @Author ftang
 * @Date 2022/3/13 11:53
 **/
import "fmt"

func makeslice()  {
	slicemake := make([]int, 2)
	fmt.Println(slicemake)
	src := []int{3,5}
	copy(slicemake, src)
	fmt.Println(slicemake)
}

func slicemake2()  {
	slicemake2 := make([]int, 2, 4)
	fmt.Println(slicemake2)
	src := []int{3,5}
	//	src := []int{3,5,6,7} 这里是不可以的，因为slicemake2的长度是2，无法把长度为4的src复制给它。
	//	只有使用append才能用到可扩容的长度4
	copy(slicemake2, src)
	slice3 := append(slicemake2,23)
	fmt.Println(slice3)

}

func main() {
	makeslice()
	slicemake2()
}

/*
使用make()函数定义，make([]type,len)
指定容量make([]type, len, cap)，cap为可变参数
容量的用处是在你用 append 扩展长度时，如果新的长度小于容量，不会更换底层数组，否则，go 会新申请一个底层数组，拷贝这边的值过去，把原来的数组丢掉。
 */