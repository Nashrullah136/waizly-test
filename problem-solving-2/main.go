package main

import "fmt"

func slicesCountCriteriaFunc[T any](data []T, criteriaFunc func(T) bool) int {
	count := 0
	for _, datum := range data {
		if criteriaFunc(datum) {
			count += 1
		}
	}
	return count
}

func fractionOfNegative(data []int) float64 {
	total := float64(len(data))
	count := float64(slicesCountCriteriaFunc(data, func(i int) bool {
		return i < 0
	}))
	return count / total
}

func fractionOfZero(data []int) float64 {
	total := float64(len(data))
	count := float64(slicesCountCriteriaFunc(data, func(i int) bool {
		return i == 0
	}))
	return count / total
}

func fractionOfPositive(data []int) float64 {
	total := float64(len(data))
	count := float64(slicesCountCriteriaFunc(data, func(i int) bool {
		return i > 0
	}))
	return count / total
}

func inputTillValid(format string, a ...any) {
	var err error
	for {
		_, err = fmt.Scanf(format, a...)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		return
	}
}

func main() {
	var n int
	inputTillValid("%d\n", &n)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		inputTillValid("%d", &data[i])
	}
	fmt.Printf("%.6f ", fractionOfPositive(data))
	fmt.Printf("%.6f ", fractionOfNegative(data))
	fmt.Printf("%.6f", fractionOfZero(data))
}
