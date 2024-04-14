func convert(s string, numRows int) string {
	rows := make([]string, numRows)
	rowIdx := 0
	dir := 1
	for _, c := range s {
		rows[rowIdx] += string(c)
		rowIdx = (rowIdx + dir) % numRows
		if rowIdx == numRows-1 || rowIdx == 0 {
			dir *= -1
		}
	}

	var result string
	for _, row := range rows {
		result += row
	}
	return result
}