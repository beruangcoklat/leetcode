func average(salary []int) float64 {
	total := salary[0]
	min := salary[0]
	max := salary[0]
	length := len(salary)

	for i := 1; i < length; i++ {
		total += salary[i]
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
	}

	return float64(total-max-min) / float64(length-2)
}
