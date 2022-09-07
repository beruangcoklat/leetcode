func countPrimes(n int) int {
	arr := make([]bool, n)
	for i := 0; i < n; i++ {
		if i <= 1 {
			continue
		}
		arr[i] = true
	}

	for i := 2; i*i < n; i++ {
		for j := i * i; j < n; j += i {
			arr[j] = false
		}
	}

	ret := 0
	for _, a := range arr {
		if a {
			ret++
		}
	}
	return ret
}
