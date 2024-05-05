
func getSatuan(param int) string {
	switch param {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	}
	return ""
}

func getPuluhan(param int) string {
	switch param {
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	}
	return ""
}

func getRatusan(param int) string {
	switch param {
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	}
	return ""
}

func getRibuan(param int) string {
	switch param {
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	}
	return ""
}

func intToRoman(num int) string {
	var ret string
	x := 1
	for i := 0; i < 4; i++ {
		r := num % (10 * x) / (1 * x)
		x *= 10

		var rr string
		switch i {
		case 0:
			rr = getSatuan(r)
		case 1:
			rr = getPuluhan(r)
		case 2:
			rr = getRatusan(r)
		case 3:
			rr = getRibuan(r)
		}

		ret = rr + ret
	}

	return ret
}
