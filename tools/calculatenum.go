package tools

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
	var min int64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// CalcalateMinFloat64 calculate min
func CalcalateMinFloat64(nums []float64) float64 {
	var min float64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}
