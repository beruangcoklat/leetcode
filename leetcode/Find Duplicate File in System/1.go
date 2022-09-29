func getFilenameAndContent(str string) (string, string) {
	filename := ""
	content := ""
	foundBracket := false

	for _, c := range str {
		if c == '(' {
			foundBracket = true
			continue
		}

		if c == ')' {
			break
		}

		if !foundBracket {
			filename += string(c)
		} else {
			content += string(c)
		}
	}
	return filename, content
}

func findDuplicate(paths []string) [][]string {
	m := map[string][]string{}
	for _, path := range paths {
		splits := strings.Split(path, " ")
		folder := ""
		startIdx := 0

		if !strings.Contains(splits[0], "(") {
			folder = splits[0]
			startIdx = 1
		}

		for i := startIdx; i < len(splits); i++ {
			filename, content := getFilenameAndContent(splits[i])
			p := fmt.Sprintf("%s/%s", folder, filename)
			m[content] = append(m[content], p)
		}
	}

	ret := [][]string{}
	for _, v := range m {
		if len(v) < 2 {
			continue
		}
		ret = append(ret, v)
	}
	return ret
}
