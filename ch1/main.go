package main

import (
	"fmt"
	"math"
)

var P0 = []int{
	0, 1, 1, 1, 0,
	1, 0, 0, 0, 1,
	1, 0, 0, 0, 1,
	1, 0, 0, 0, 1,
	0, 1, 1, 1, 0,
}

var P1 = []int{
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
	0, 0, 1, 0, 0,
}
var P2 = []int{
	0, 1, 1, 1, 1,
	1, 0, 0, 1, 0,
	0, 0, 1, 0, 0,
	0, 1, 0, 0, 0,
	1, 1, 1, 1, 1,
}
var P3 = []int{
	0, 1, 1, 1, 0,
	1, 0, 0, 0, 1,
	0, 0, 1, 1, 0,
	1, 0, 0, 0, 1,
	0, 1, 1, 1, 0,
}
var P4 = []int{
	0, 0, 1, 0, 0,
	0, 1, 0, 0, 0,
	1, 0, 0, 1, 0,
	1, 1, 1, 1, 1,
	0, 0, 0, 1, 0,
}
var P5 = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 0, 0,
	1, 1, 1, 1, 0,
	0, 0, 0, 0, 1,
	1, 1, 1, 1, 0,
}
var P6 = []int{
	0, 0, 1, 1, 0,
	0, 1, 0, 0, 0,
	1, 1, 1, 1, 0,
	1, 0, 0, 0, 1,
	0, 1, 1, 1, 0,
}
var P7 = []int{
	1, 1, 1, 1, 1,
	1, 0, 0, 1, 0,
	0, 0, 1, 0, 0,
	0, 1, 0, 0, 0,
	1, 0, 0, 0, 0,
}
var P8 = []int{
	0, 1, 1, 1, 0,
	1, 0, 0, 0, 1,
	0, 1, 1, 1, 0,
	1, 0, 0, 0, 1,
	0, 1, 1, 1, 0,
}
var P9 = []int{
	0, 1, 1, 1, 0,
	1, 0, 0, 0, 1,
	0, 1, 1, 1, 1,
	0, 0, 0, 1, 0,
	0, 0, 1, 0, 0,
}

func main() {
	x1 := []int{
		0, 1, 1, 1, 0,
		0, 1, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
		1, 1, 1, 1, 1,
	}
	x2 := []int{
		0, 1, 0, 0, 0,
		0, 1, 0, 1, 0,
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
	}
	x3 := []int{
		0, 1, 1, 1, 1,
		1, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
	}
	x4 := []int{
		0, 1, 0, 1, 0,
		1, 1, 1, 1, 1,
		0, 1, 0, 1, 0,
		0, 1, 1, 1, 0,
		1, 1, 0, 1, 1,
	}
	calcVector1(x1, "x1")
	calcVector1(x2, "x2")
	calcVector1(x3, "x3")
	calcVector1(x4, "x4")
	calcVector2(x1, "x1")
	calcVector2(x2, "x2")
	calcVector2(x3, "x3")
	calcVector2(x4, "x4")
}
func calcVector2(x []int, name string) {
	fmt.Println("name is ", name)
	boxSize := 5
	list := [][]int{P0, P1, P2, P3, P4, P5, P6, P7, P8, P9}

	for tempIdx, temp := range list {
		sum := 0
		tempSum := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
		valSum := [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
		for i := 0; i < len(x); i++ {
			w := i % boxSize
			h := i / boxSize
			tempSum[0][h] = tempSum[0][h] + temp[i]
			tempSum[1][w] = tempSum[1][w] + temp[i]
			valSum[0][h] = valSum[0][h] + x[i]
			valSum[1][w] = valSum[1][w] + x[i]
		}
		for i := range tempSum {
			ts := tempSum[i]
			vs := valSum[i]
			for idx := range ts {
				val := ts[idx] - vs[idx]
				vector := int(math.Abs(float64(val)))
				sum = sum + (vector * vector)
			}

		}
		fmt.Printf("%02d sum=%d\n", tempIdx+1, sum)
	}

}
func calcVector1(x []int, name string) {
	fmt.Println("name is ", name)
	list := [][]int{P0, P1, P2, P3, P4, P5, P6, P7, P8, P9}
	for _, temp := range list {
		sum := 0
		for i := 0; i < len(x); i++ {
			val := temp[i] - x[i]
			vector := int(math.Abs(float64(val)))
			sum = sum + vector
		}
		fmt.Printf("sum=%d\n", sum)
	}

}
