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
	//两种不同的字符串表示方法
	var website = "\thttps://www.callicoder.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`

	fmt.Println(website, siteDescription)

	// 类型转换
	var m int64 = 4
	var n int = int(m)
	var c float64 = 6.5
	var result = float64(n) + c
	fmt.Println(result)

	//常量类型
	const myFavLanguage = "python"
	const sunRiseintheEast = true
	fmt.Println(myFavLanguage, sunRiseintheEast)
	var age = 18
	if age >= 18 {
		fmt.Println("you are eligible to vote! ")
	} else {
		fmt.Println("you are not eligible to vote! ")
	}
}
