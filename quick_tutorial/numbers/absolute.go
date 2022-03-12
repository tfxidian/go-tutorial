package numbers

import (
	"math"
)

func absolute(num int)bool{
	if num <0{
		return 0-num
	}
	return num
}