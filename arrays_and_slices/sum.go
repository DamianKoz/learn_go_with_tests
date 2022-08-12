package main

func Sum(nums []int) (sum int) {
	for _, number := range nums {
		sum += number
	}
	return sum
}
