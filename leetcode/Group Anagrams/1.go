func sortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] > bytes[j]
	})
	return string(bytes)
}

func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)

		_, ok := groupMap[sortedStr]
		if !ok {
			groupMap[sortedStr] = []string{str}
		} else {
			groupMap[sortedStr] = append(groupMap[sortedStr], str)
		}
	}

	groups := [][]string{}
	for _, list := range groupMap {
		groups = append(groups, list)
	}

	return groups
}