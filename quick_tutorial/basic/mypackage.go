// Go file
/***********************************************
# Copyright (c) 2022, XiAn
# All rights reserved.
#
# @Filename: package.go
# @Version：V1.0
# @Author: ftang - tfxidian@163.com
# @Description: ---using exsit go package
# @Create Time: 2022-03-12 13:14:30
# @Last Modified: 2022-03-12 13:14:30
***********************************************/

package main

import(
    "fmt"
    "github.com/tfxidian/go-tutorial/quick_tutorial/numbers"
)


func main(){
    fmt.Println(numbers.Absolute(-4))
    fmt.Println(numbers.IsPrime(19))
}
