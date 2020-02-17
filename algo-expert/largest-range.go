package program

func LargestRange(array []int) []int {
	dict := make(map[int]struct{})
	min := 0
	for i := 0; i < len(array); i++ {
		if i == 0 {
			min = array[i]
		} else if array[i] < min {
			min = array[i]
		}
		dict[array[i]] = struct{}{}
	}

	max := 0
	maxStart, maxEnd := 0, 0
	for k, _ := range dict {
		curr := 1
		next := k + 1
		for true {
			_, ok := dict[next]
			if ok {
				curr += 1
				if curr > max {
					maxStart = k
					maxEnd = next
					max = curr
				}
				next += 1
				continue
			}

			if max < curr {
				max = curr
				maxStart = k
				maxEnd = next - 1
			}

			break
		}
	}

	return []int{maxStart, maxEnd}
}
