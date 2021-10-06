func moveZeroes(nums []int) {
	notZero := []int{}
	zeroCount := 0
	for _, n := range nums {
		if n == 0 {
			zeroCount++
		} else {
			notZero = append(notZero, n)
		}
	}

	for i := 0; i < len(notZero); i++ {
		nums[i] = notZero[i]
	}

	for i := 0; i < zeroCount; i++ {
		nums[i+len(notZero)] = 0
	}
}
