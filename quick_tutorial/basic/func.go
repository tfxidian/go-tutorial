package main

import(
	"fmt"
)

func add(x int, y int) int{
    return (x+y)/2
}


func FullNames(name1 string, name2 string)(string, string){
    newName1 := "Zhang"+name1
    newName2 := "zhang"+name2
    return newName1, newName2
}


func main(){
    fmt.Println("hello func")
    result := add(3,5)
    fmt.Println(result)
    name1, name2 := FullNames("ming", "liang")
    fmt.Println(name1, name2)
}
