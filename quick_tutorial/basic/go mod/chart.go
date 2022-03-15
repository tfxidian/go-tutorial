package main

/**
 * @Description
 * @Author ftang
 * @Date 2022/3/15 8:37
 **/

import (
	"github.com/wcharczuk/go-chart"
	"os"
)

func main()  {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: []float64{1.0, 2.0, 3.0, 4.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0},
			},
		},
	}
	f, _ :=os.Create("output.PNG")
	defer  f.Close()
	graph.Render(chart.PNG, f)
}
