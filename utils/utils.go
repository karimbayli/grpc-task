package utils

import (
	"fmt"
	"sort"
)

func CalcMedian(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty employee list")
	}
	sort.Float64s(numbers) // sort the numbers

	mNumber := len(numbers) / 2

	if len(numbers)%2 != 0 {
		return numbers[mNumber], nil
	}

	return (numbers[mNumber-1] + numbers[mNumber]) / 2, nil
}
