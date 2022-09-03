package main

import "sort"

func SortedSquaredArray(array []int) []int {
	for i := 0; i < len(array); i++ {
		array[i] *= array[i]
	}
	sort.SliceStable(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	return array
}
