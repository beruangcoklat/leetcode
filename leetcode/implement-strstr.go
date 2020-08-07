// https://leetcode.com/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle > lenHaystack {
		return -1
	}

	if lenNeedle == lenHaystack {
		if needle == haystack {
			return 0
		}
		return -1
	}

	for i := 0; i < lenHaystack-lenNeedle+1; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}

	return -1
}