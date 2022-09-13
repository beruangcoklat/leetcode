func countPrimes(n int) int {
	arr := make([]bool, n)
	for i := 0; i < n; i++ {
		arr[i] = true
	}

	for i := 2; i < n; i++ {
		for j := i * 2; j < n; j += i {
			arr[j] = false
		}
	}

	if len(arr) > 0 {
		arr[0] = false
	}
	if len(arr) > 1 {
		arr[1] = false
	}

	ret := 0
	for _, a := range arr {
		if a {
			ret++
		}
	}
	return ret
}
