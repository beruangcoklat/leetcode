var result [][]int

func solve(wordIdx int, words []string, score []int, quotaChar []int, resultWordIdx []int) {
	if wordIdx == len(words) {
		result = append(result, resultWordIdx)
		return
	}

	valid, quotaCharClone := validWord(words[wordIdx], quotaChar)
	if valid {
		resultWordIdxClone := make([]int, len(resultWordIdx)+1)
		copy(resultWordIdxClone, resultWordIdx)
		resultWordIdxClone[len(resultWordIdx)] = wordIdx
		solve(wordIdx+1, words, score, quotaCharClone, resultWordIdxClone)
	}

	solve(wordIdx+1, words, score, quotaChar, resultWordIdx)
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	quotaChar := make([]int, 26)
	for _, char := range letters {
		quotaChar[getIdxFromByte(char)]++
	}

	result = make([][]int, 0)
	solve(0, words, score, quotaChar, make([]int, 0))

	maxScore := 0
	for _, r := range result {
		currScore := 0
		for _, idx := range r {
			currScore += calculateScore(words[idx], score)
		}

		if currScore > maxScore {
			maxScore = currScore
		}
	}
	return maxScore
}

func validWord(word string, quotaChar []int) (bool, []int) {
	quotaCharClone := make([]int, 26)
	copy(quotaCharClone, quotaChar)

	valid := true
	for _, char := range word {
		idx := getIdxFromByte(byte(char))
		quotaCharClone[idx]--
		if quotaCharClone[idx] < 0 {
			valid = false
			break
		}
	}

	if valid {
		return valid, quotaCharClone
	}
	return valid, quotaChar
}

func calculateScore(word string, score []int) int {
	currScore := 0
	for _, char := range word {
		idx := getIdxFromByte(byte(char))
		currScore += score[idx]
	}
	return currScore
}

func getIdxFromByte(a byte) int {
	return int(a) - 'a'
}
