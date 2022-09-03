package main

func IsValidSubsequence(array []int, sequence []int) bool {
	seqIdx := 0
	for _, a := range array {
		if a == sequence[seqIdx] {
			seqIdx++
		}
		if seqIdx == len(sequence) {
			return true
		}
	}
	return false
}
