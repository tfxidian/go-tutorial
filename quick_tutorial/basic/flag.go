package main

import (
	"flag"
	"fmt"
)

var cliName = flag.String("name", "John", "Input Your Name")
var cliAge = flag.Int("age", 22, "Input Your Age")
var cliGender = flag.String("gender", "male", "Input Your Gender")

func main() {

	flag.Parse()
	// 输出命令行参数
	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	//fmt.Printf("period=", *cliPeriod)
	// flag.Args() 函数返回没有被解析的命令行参数
	// func NArg() 函数返回没有被解析的命令行参数的个数
	//fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
}
