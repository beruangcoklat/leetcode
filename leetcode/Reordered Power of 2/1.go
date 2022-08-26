var result bool

func permutation(m map[string]struct{}, s, curr string, allVisited, visited int) {
	if result {
		return
	}

	if allVisited == visited {
		_, ok := m[curr]
		if ok {
			result = true
		}
		return
	}

	for i := 0; i < len(s); i++ {
		if (visited & (1 << i)) != 0 {
			continue
		}

		permutation(
			m,
			s,
			curr+string(s[i]),
			allVisited,
			visited|(1<<i))
	}
}

func reorderedPowerOf2(n int) bool {
	var a int64 = 1
	m := map[string]struct{}{}
	for a <= math.MaxInt32 {
		m[strconv.Itoa(int(a))] = struct{}{}
		a *= 2
	}
	result = false

	s := strconv.Itoa(n)
	allVisited := (1 << len(s)) - 1
	permutation(m, s, "", allVisited, 0)
	return result
}
