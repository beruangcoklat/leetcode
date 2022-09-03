package main

func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]struct{})

	for _, a := range array {
		need := target - a
		_, ok := m[need]
		if ok {
			return []int{a, need}
		}
		m[a] = struct{}{}
	}

	return []int{}
}
