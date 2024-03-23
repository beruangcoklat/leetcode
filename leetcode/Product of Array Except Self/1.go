func productExceptSelf(nums []int) []int {
	size := len(nums)
	multiplyFromLeft := make([]int, size)
	multiplyFromRight := make([]int, size)
	result := make([]int, size)

	multiplyFromLeft[0] = nums[0]
	for i := 1; i < size; i++ {
		multiplyFromLeft[i] = nums[i] * multiplyFromLeft[i-1]
	}

	multiplyFromRight[size-1] = nums[size-1]
	for i := size - 2; i >= 0; i-- {
		multiplyFromRight[i] = nums[i] * multiplyFromRight[i+1]
	}

	for i := 0; i < size; i++ {
		result[i] = 1
		if i-1 >= 0 {
			result[i] *= multiplyFromLeft[i-1]
		}
		if i+1 < size {
			result[i] *= multiplyFromRight[i+1]
		}
	}

	return result
}
