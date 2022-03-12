// Go file
/***********************************************
# Copyright (c) 2022, XiAn
# All rights reserved.
#
# @Filename: array.go
# @Version：V1.0
# @Author: ftang - tfxidian@163.com
# @Description: ---
# @Create Time: 2022-03-12 12:23:40
# @Last Modified: 2022-03-12 12:23:40
***********************************************/

package main

import "fmt"

func arraydemo(){
    var array1[4] int
    array1[0]= 1
    array1[1]= 2
    array1[2]= 3
    array1[3]= 4
    fmt.Println(len(array1))
    fmt.Println(array1)
}

func arraydemo2(){
    array2 := [3]string{"num1", "num2", "num3"}
    fmt.Println(array2)
}

func arraydemo3(){
    array3 := [2][2]string{
        {"cpp", "c"},
        {"java", "go"}, //note that here is a , 
    }
    fmt.Println(array3)

}

func arraydemo4(){
    array4 := [...]int{1,3,5,6,8,2}
    fmt.Println(array4)
    fmt.Println(len(array4))
}
func iter_array(){
    array0 := []int{1,2,3,4,5}
    for x :=0; x<5; x++{
        fmt.Println(array0[x])
    }
}


func main(){
    arraydemo()
    arraydemo2()
    arraydemo3()
    iter_array()
}
