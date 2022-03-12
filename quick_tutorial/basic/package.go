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
    "math"
    "math/rand"
    "time"
    "github.com/callicoder/golang-tutorials/07-packages/numbers"
	"github.com/callicoder/golang-tutorials/07-packages/strings"
	"github.com/callicoder/golang-tutorials/07-packages/strings/greeting" // Importing a nested package
    str "strings"
)


func main(){
    fmt.Println(math.Max(100,200))
    timenow := time.Now().Unix()
    fmt.Print(timenow)
    rand.Seed(timenow)
    fmt.Println(rand.Intn(100))

    fmt.Println(numbers.IsPrime(19))
    fmt.Println(greeting.WelcomeText)
    fmt.Println(strings.Reverse("callicoder"))
    fmt.Println(str.Count("go is awesome", "go"))
}
