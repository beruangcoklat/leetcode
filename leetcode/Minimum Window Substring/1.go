func same(tMap, currMap map[string]int) bool {
	for k, v := range tMap {
		if currMap[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	tMap := make(map[string]int)
	for _, c := range t {
		tMap[string(c)]++
	}

	size := len(s)
	leftIdx, rightIdx := 0, 0
	currMap := make(map[string]int)

	minResultFound := false
	currMap[string(s[0])]++
	minResult := 0
	leftResult, rightResult := 0, 0

	for {
		if same(tMap, currMap) {
			length := rightIdx - leftIdx + 1
			if length < minResult || !minResultFound {
				minResultFound = true
				minResult = length
				leftResult = leftIdx
				rightResult = rightIdx
			}

			currMap[string(s[leftIdx])]--
			leftIdx++
		} else {
			rightIdx++
			if rightIdx >= size {
				break
			}
			currMap[string(s[rightIdx])]++
		}
	}

	if !minResultFound {
		return ""
	}
	return s[leftResult : rightResult+1]
}
