func minMovesToSeat(seats []int, students []int) int {
	sort.SliceStable(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})
	sort.SliceStable(students, func(i, j int) bool {
		return students[i] < students[j]
	})
	ret := 0
	for i := 0; i < len(seats); i++ {
		ret += abs(seats[i] - students[i])
	}
	return ret
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
