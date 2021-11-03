func pow(a, b uint32) uint32 {
	if b == 0 {
		return 1
	}
	ret := uint32(1)
	for i := uint32(0); i < b; i++ {
		ret *= a
	}
	return ret
}

func reverseBits(num uint32) uint32 {
	arr := make([]uint32, 32)
	idx := 0
	for num > 0 {
		arr[idx] = num % 2
		idx++
		num = num >> 1
	}

	ret := uint32(0)
	for i := len(arr) - 1; i >= 0; i-- {
		idx := len(arr) - 1 - i
		if arr[idx] == 0 {
			continue
		}
		ret += pow(2, uint32(i))
	}

	return ret
}
