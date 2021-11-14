type data struct {
	val   int
	valid bool
}

var memo map[int]data

func solve(idx int, nums []int) data {
	if idx == len(nums)-1 {
		return data{valid: true}
	}

	if val, ok := memo[idx]; ok {
		return val
	}

	if nums[idx] == 0 {
		return data{valid: false}
	}

	min := 0
	valid := false
	for i := 1; i <= nums[idx]; i++ {
		if idx+i >= len(nums) {
			break
		}

		curr := solve(idx+i, nums)
		if !curr.valid {
			continue
		}

		if !valid || curr.val+1 < min {
			valid = true
			min = curr.val + 1
		}
	}

	if !valid {
		return data{valid: false}
	}

	res := data{
		val:   min,
		valid: true,
	}
	memo[idx] = res
	return res
}

func jump(nums []int) int {
	memo = make(map[int]data)
	return solve(0, nums).val
}
