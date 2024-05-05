var m = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	ret := ""
	for num > 0 {
		for i := 0; i < len(nums); i++ {
			if num >= nums[i] {
				ret += m[nums[i]]
				num -= nums[i]
				break
			}
		}
	}
	return ret
}
