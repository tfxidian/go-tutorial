// Go file
/***********************************************
# Copyright (c) 2022, XiAn
# All rights reserved.
#
# @Filename: hello.go
# @Version：V1.0
# @Author: ftang - tfxidian@163.com
# @Description: ---
# @Create Time: 2022-03-12 14:23:09
# @Last Modified: 2022-03-12 14:23:09
***********************************************/
package main
import "fmt"
import "flag"
import "github.com/golang/glog"
func main() {
    flag.Parse()
    fmt.Println("Hello World!")
    glog.Info("glog Hello World!")
    glog.Flush()
} 
