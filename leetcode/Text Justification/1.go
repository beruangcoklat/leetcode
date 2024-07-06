func fullJustify(words []string, maxWidth int) []string {
	result := []string{""}
	idx := 0
	for _, word := range words {
		if (result[idx] == "" && len(word)+len(result[idx]) <= maxWidth) || (result[idx] != "" && len(word)+len(result[idx]) < maxWidth) {
			if result[idx] == "" {
				result[idx] = word
			} else {
				result[idx] += " " + word
			}
		} else {
			idx++
			result = append(result, "")
			if result[idx] == "" {
				result[idx] = word
			} else {
				result[idx] += " " + word
			}
		}
	}

	for i, line := range result {
		if i == len(result)-1 {
			result[i] += getSpace(maxWidth - len(line))
			continue
		}

		words := strings.Split(line, " ")
		if len(words) == 1 {
			result[i] = words[0] + getSpace(maxWidth-len(words[0]))
			continue
		}

		spaceCount := len(words) - 1
		extraSpace := maxWidth - len(line)

		space := (extraSpace / spaceCount) + 1

		result[i] = words[0]
		additionalSpace := extraSpace % spaceCount
		for j := 1; j < len(words); j++ {
			if additionalSpace > 0 {
				additionalSpace--
				result[i] += getSpace(space+1) + words[j]
			} else {
				result[i] += getSpace(space) + words[j]
			}

		}
	}

	return result
}

func getSpace(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(" ")
	}
	return b.String()
}
