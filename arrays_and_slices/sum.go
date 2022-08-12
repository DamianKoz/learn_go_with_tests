package main

import "fmt"

func Sum(nums []int) (sum int) {
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, slice := range numbersToSum {
		tempSum := 0
		for _, number := range slice {
			tempSum += number
		}
		result = append(result, tempSum)
	}
	return result
}

func main() {
	fmt.Print(SumAll([]int{1, 2}, []int{2, 3}))
}
