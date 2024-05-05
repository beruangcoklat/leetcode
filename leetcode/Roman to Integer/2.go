var m = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func calculate(a string) int {
	value := m[a[len(a)-1]]
	for i := len(a) - 2; i >= 0; i-- {
		value -= m[a[i]]
	}
	return value
}

func romanToInt(s string) int {
	size := len(s)
	ret := 0
	for i := 0; i < size; i++ {
		if i+1 >= size || m[s[i]] >= m[s[i+1]] {
			ret += calculate(string(s[i]))
			continue
		}

		str := string(s[i])
		idx := i
		for idx+1 < size && m[s[idx+1]] > m[s[idx]] {
			str += string(s[idx+1])
			idx++
		}
		ret += calculate(str)
		i = idx
	}

	return ret
}
