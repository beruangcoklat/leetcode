func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isE(c byte) bool {
	return c == 'e' || c == 'E'
}

func isPlusOrMinus(c byte) bool {
	return c == '+' || c == '-'
}

func isNumber(s string) bool {
	validCharacter := map[byte]struct{}{
		'0': {},
		'1': {},
		'2': {},
		'3': {},
		'4': {},
		'5': {},
		'6': {},
		'7': {},
		'8': {},
		'9': {},
		'e': {},
		'E': {},
		'+': {},
		'-': {},
		'.': {},
	}

	eAppear := false
	dotAppear := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := validCharacter[c]; !ok {
			return false
		}

		if eAppear && (c == '.' || isE(c)) {
			return false
		}

		if dotAppear && c == '.' {
			return false
		}

		if isE(c) {
			eAppear = true
		} else if c == '.' {
			dotAppear = true
		}

		if isPlusOrMinus(c) {
			if i > 0 && !isE(s[i-1]) {
				return false
			}
			if i+1 >= len(s) || (!isDigit(s[i+1]) && s[i+1] != '.') {
				return false
			}
		}

		if isE(c) && (i == 0 || i+1 >= len(s) || (!isDigit(s[i+1]) && !isPlusOrMinus(s[i+1]))) {
			return false
		}

		if c == '.' {
			if i-1 >= 0 && isDigit(s[i-1]) {
				continue
			}
			if i+1 < len(s) && isDigit(s[i+1]) {
				continue
			}
			return false
		}
	}

	return true
}
