func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[rune]int)
	for _, c := range magazine {
		magazineMap[c]++
	}
	for _, c := range ransomNote {
		magazineMap[c]--
		if magazineMap[c] < 0 {
			return false
		}
	}
	return true
}
