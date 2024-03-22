func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}

	leftSum, rightSum := 1, n
	leftIdx, rightIdx := 2, n-1

	for leftIdx < rightIdx {
		// fmt.Println(leftIdx, leftSum, rightIdx, rightSum)

		if leftSum <= rightSum {
			leftSum += leftIdx
			leftIdx++
		} else {
			rightSum += rightIdx
			rightIdx--
		}
	}

	if leftSum == rightSum {
		return leftIdx
	}

	return -1
}