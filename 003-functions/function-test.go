package main

import (
	"errors"
	"fmt"
)

//basic function
func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

//两个返回值
func getStockPriceChange(prevPrice float64, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	fmt.Println(percentChange)
	return change, percentChange
}

//多个返回值
func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("previus price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func main() {
	//基本函数测试
	x := 5.75
	y := 6.25
	result := avg(x, y)
	fmt.Printf("average of %.2f and %.2f is %2.f\n", x, y, result)

	//两个返回值测试
	prevStockPrice := 50000.0
	currentStockPrice := 100000.0
	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)
	fmt.Printf("change is %2.f, percentChange is %f%%\n", change, percentChange)
	if change < 0 {
		fmt.Printf("The stock price decreased .\n")
	} else {
		fmt.Println("the stock price incresed")
	}
	//多个返回值测试

	change2, percentChange2, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)
	fmt.Printf("change2 is %f, percentChange2 is %f\n", change2, percentChange2)
	fmt.Printf("type of err %T\n", err)
	if err != nil {
		fmt.Println("sorry, there was an error", err)
	} else {
		if change < 0 {
			fmt.Println("the stock price decreased...")
		} else {
			fmt.Println("the stock price increased....")
		}
	}

}
