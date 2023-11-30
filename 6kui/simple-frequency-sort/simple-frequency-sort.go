package kata

import (
	"sort"
)

type freq struct {
	value int
	count int
}

func Solve(arr []int) []int {
	result := []int{}
	frequencyMap := make(map[int]int)
	for _, num := range arr {
		frequencyMap[num]++
	}
	freqSlice := make([]freq, 0, len(frequencyMap))
	for value, count := range frequencyMap {
		freqSlice = append(freqSlice, freq{value, count})
	}
	sort.Slice(freqSlice, func(i, j int) bool {
		if freqSlice[i].count == freqSlice[j].count {
			return freqSlice[i].value < freqSlice[j].value
		}
		return freqSlice[i].count > freqSlice[j].count
	})
	for _, f := range freqSlice {
		for i := 0; i < f.count; i++ {
			result = append(result, f.value)
		}
	}
	return result
}

/*
In this kata, you will sort elements in an array by decreasing frequency of elements.
If two elements have the same frequency, sort them by increasing value.

solve([2,3,5,3,7,9,5,3,7]) = [3,3,3,5,5,7,7,2,9]
-- We sort by highest frequency to lowest frequency.
-- If two elements have same frequency, we sort by increasing
*/
