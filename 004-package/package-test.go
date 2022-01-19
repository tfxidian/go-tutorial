package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//使用数学package
	fmt.Println(math.Max(23, 45))
	fmt.Println(math.Sqrt(169))
	//使用时间package
	epoch := time.Now().Unix()
	fmt.Println(epoch)
	//使用 rand package
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))

}
