package main

import "math"

// Calculate the mean of a list of numbers
func mean(numbers []float64) float64 {
	return sum(numbers) / float64(len(numbers))
}

// Calculate the standard deviation of a list of numbers
func stdev(numbers []float64) float64 {
	var sd float64
	for j := 0; j < len(numbers); j++ {
		// The use of Pow math function func Pow(x, y float64) float64
		sd += math.Pow(numbers[j]-mean(numbers), 2)
	}
	// The use of Sqrt math function func Sqrt(x float64) float64
	sd = math.Sqrt(sd / float64(len(numbers)))
	return sd //sqrt(variance)
}

func sum(arr []float64) float64 {
	var res float64
	res = 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}
