
func get(s string, i int) string {
	if i >= len(s) {
		return ""
	}
	return s[i : i+1]
}

func romanToInt(s string) int {
	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var ret int
	for i := 0; i < len(s); i++ {
		curr := get(s, i)
		next := get(s, i+1)

		currNum := dict[curr]

		if next == "" {
			ret += currNum
			continue
		}

		nextNum := dict[next]

		if currNum >= nextNum {
			ret += currNum
			continue
		}

		ret += (nextNum - currNum)
		i++
	}

	return ret
}
