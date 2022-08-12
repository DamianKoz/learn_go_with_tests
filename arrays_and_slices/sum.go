package main

func Sum(nums []int) (sum int) {
	for _, i := range nums {
		sum += i
	}
	return sum
}
