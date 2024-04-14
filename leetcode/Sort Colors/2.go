func merge(left, mid, right int, arr []int) {
	arr1 := []int{}
	arr2 := []int{}
	leftIdx := 0
	rightIdx := 0
	idx := left

	for i := left; i <= mid; i++ {
		arr1 = append(arr1, arr[i])
	}

	for i := mid + 1; i <= right; i++ {
		arr2 = append(arr2, arr[i])
	}

	for leftIdx < len(arr1) && rightIdx < len(arr2) {
		var min int
		if arr1[leftIdx] <= arr2[rightIdx] {
			min = arr1[leftIdx]
			leftIdx++
		} else {
			min = arr2[rightIdx]
			rightIdx++
		}

		arr[idx] = min
		idx++
	}

	for leftIdx < len(arr1) {
		arr[idx] = arr1[leftIdx]
		leftIdx++
		idx++
	}

	for rightIdx < len(arr2) {
		arr[idx] = arr2[rightIdx]
		rightIdx++
		idx++
	}

}

func mergeSort(left, mid, right int, arr []int) {
	if left >= right {
		return
	}
	mergeSort(left, (left+mid)/2, mid, arr)
	mergeSort(mid+1, (mid+1+right)/2, right, arr)
	merge(left, mid, right, arr)
}

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	mid := (right - left) / 2
	mergeSort(left, mid, right, nums)
	fmt.Println(nums)
}