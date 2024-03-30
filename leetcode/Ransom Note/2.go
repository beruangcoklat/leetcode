func canConstruct(ransomNote string, magazine string) bool {
	ransomNoteMap := make(map[rune]int)
	for _, c := range ransomNote {
		ransomNoteMap[c]++
	}
	for _, c := range magazine {
		ransomNoteMap[c]--
	}

	for _, v := range ransomNoteMap {
		if v > 0 {
			return false
		}
	}
	return true
}
