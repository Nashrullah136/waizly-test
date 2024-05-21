package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sum(data ...int64) int64 {
	total := int64(0)
	for _, datum := range data {
		total += datum
	}
	return total
}

func minOfFour(data []int64) int64 {
	maxData := slices.MaxFunc(data, func(a, b int64) int {
		return cmp.Compare(a, b)
	})
	return sum(data...) - maxData
}

func maxOfFour(data []int64) int64 {
	minData := slices.MinFunc(data, func(a, b int64) int {
		return cmp.Compare(a, b)
	})
	return sum(data...) - minData
}

func main() {
	data := make([]int64, 5)
	var err error
	for {
		_, err = fmt.Scanf("%d %d %d %d %d\n", &data[0], &data[1], &data[2], &data[3], &data[4])
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		break
	}
	fmt.Printf("%d %d", minOfFour(data), maxOfFour(data))
}
