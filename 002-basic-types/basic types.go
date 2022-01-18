package main

import "fmt"

func main() {
	//数据类型
	var myInt8 int8 = 97
	var myInt = 1200
	var myUnit uint = 500
	var myHexNumber = 0xff
	var myOctalNum = 034
	var myfloat32 float32 = 4.5

	fmt.Printf("%d, %d, %d,  %#x, %#o %f \n", myInt8, myInt, myUnit, myHexNumber, myOctalNum, myfloat32)
    
}
