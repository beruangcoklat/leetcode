func numberOfWeakCharacters(properties [][]int) int {
	sort.SliceStable(properties, func(i, j int) bool {
		if properties[i][0] < properties[j][0] {
			return true
		}
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return false
	})

	firstAttack := properties[0][0]
	attackGroupIdx := []int{0}
	attackValueToGroupIdx := map[int]int{
		firstAttack: 0,
	}

	for i := 1; i < len(properties); i++ {
		prevAttack := properties[i-1][0]
		attack := properties[i][0]
		if prevAttack != attack {
			attackGroupIdx = append(attackGroupIdx, i)
			attackValueToGroupIdx[attack] = len(attackGroupIdx) - 1
		}
	}

	maxFromRight := make([]int, len(attackGroupIdx))

	for i := len(attackGroupIdx) - 1; i >= 0; i-- {
		defense := properties[attackGroupIdx[i]][1]
		if i == len(attackGroupIdx)-1 || defense > maxFromRight[i+1] {
			maxFromRight[i] = defense
		} else {
			maxFromRight[i] = maxFromRight[i+1]
		}
	}

	lastAttack := properties[len(properties)-1][0]
	ret := 0
	for i := 0; i < len(properties); i++ {
		attack := properties[i][0]
		if attack == lastAttack {
			break
		}

		attackIdx := attackValueToGroupIdx[attack]
		defense := properties[i][1]
		if defense < maxFromRight[attackIdx+1] {
			ret++
		}
	}
	return ret
}
