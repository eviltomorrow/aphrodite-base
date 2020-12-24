package tools

import "math"

// CalculateSumInt64 calculate sum
func CalculateSumInt64(nums []int64) int64 {
	var sum int64
	for _, num := range nums {
		sum += num
	}
	return sum
}

// CalculateSumFloat64 calculate sum
func CalculateSumFloat64(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

// CalcalateMaxInt64 calculate max
func CalcalateMaxInt64(nums []int64) int64 {
	var max int64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// CalcalateMaxFloat64 calculate max
func CalcalateMaxFloat64(nums []float64) float64 {
	var max float64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// CalcalateMinInt64 calculate min
func CalcalateMinInt64(nums []int64) int64 {
	if len(nums) == 0 {
		return 0
	}
	var min = nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

// CalcalateMinFloat64 calculate min
func CalcalateMinFloat64(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	var min = nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

var n10 = math.Pow10(2)

// Trunc2 trunc2
func Trunc2(val float64) float64 {
	return math.Trunc((val+0.5/n10)*n10) / n10
}
