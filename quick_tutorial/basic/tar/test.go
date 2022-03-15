package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

/**
 * @Description
 * @Author ftang
 * @Date 2022/3/15 10:55
 **/

func GetFiles()  {
	files, err := ioutil.ReadDir(".")
	if err != nil{
		log.Fatal(err)
	}

	for _, file := range files{
		fmt.Println(file.Name())
	}
}

func main() {
	fmt.Println("test begin")
	GetFiles()
}