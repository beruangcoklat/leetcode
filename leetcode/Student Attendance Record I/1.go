func checkRecord(s string) bool {
	absentCount := 0
	lastLateCount := 0
	for _, c := range s {
		switch c {
		case 'A':
			absentCount++
			lastLateCount = 0
		case 'L':
			lastLateCount++
		case 'P':
			lastLateCount = 0
		}

		if absentCount == 2 || lastLateCount == 3 {
			return false
		}
	}
	return true
}
