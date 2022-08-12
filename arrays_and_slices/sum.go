package main

func Sum(nums []int) (sum int) {
	for _, number := range nums {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, slice := range numbersToSum {
		sum := Sum(slice)
		result = append(result, sum)
	}
	return result
}

func SumAllTail(numbersToSum ...[]int) (result []int) {
	for _, slice := range numbersToSum {
		if len(slice) > 1 {
			sum := Sum(slice[1:])
			result = append(result, sum)
		} else {
			result = append(result, 0)
		}
	}
	return
}
