package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello, go!")

	var myStr string  = "hello str"
	var myint int = 100
	var myFloat float64 = 45.13
	fmt.Println(myStr, myint,myFloat)
	fmt.Printf("myStr:%s, myint: %d, myFloat: %f\n", myStr,myint, myFloat)
	
	var double_quotation_str = "\tmydouble_quotation_str\t\n"
	var single_backticks_str =  `\tsingle_quotation_str\t\n`
	fmt.Println(double_quotation_str, single_backticks_str)

}

