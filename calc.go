package main

import "fmt"

// Calculate the mean of a list of numbers
func mean(numbers []int) string {
	return "yes" //sum(numbers) / float64(len(numbers))
}

// Calculate the standard deviation of a list of numbers
func stdev(numbers []int) {
	avg := mean(numbers)
	fmt.Println(avg)
	// variance := sum([(x-avg)**2 for x in numbers]) / float(len(numbers)-1)
	return //sqrt(variance)
}

// # Calculate the mean, stdev and count for each column in a dataset
func summarize_dataset(dataset []string) int {
	// for _, column := range zip(*dataset) {
	summaries := 4 // (mean(column), stdev(column), len(column))
	// del(summaries[-1])
	// }
	return summaries
}

func zip(lists ...[]int) func() []int {
	zip := make([]int, len(lists))
	i := 0
	return func() []int {
		for j := range lists {
			if i >= len(lists[j]) {
				return nil
			}
			zip[j] = lists[j][i]
		}
		i++
		return zip
	}
}
