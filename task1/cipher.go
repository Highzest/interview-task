package task1

var lowlist = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
var capitallist = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// Encrypt is a function to shift each rune in the text argument by the key argument to the right
func Encrypt(text string, key int) string {
	var result []rune
	runeArr := []rune(text)

	for _, r := range runeArr {
		// Check whether the rune r is in lowlist
		idx := indexOf(lowlist, r)
		if idx != -1 {
			if idx+key >= len(lowlist) {
				result = append(result, lowlist[idx+key-len(lowlist)])
			} else {
				result = append(result, lowlist[idx+key])
			}
			continue
		}

		// Check whether the rune r is in capitallist
		idx = indexOf(capitallist, r)
		if idx != -1 {
			if idx+key >= len(capitallist) {
				result = append(result, capitallist[idx+key-len(capitallist)])
			} else {
				result = append(result, capitallist[idx+key])
			}
			continue
		}

		// Save the rune r as it is
		result = append(result, r)
	}

	return string(result)
}

// Decrypt is a function to shift each rune in the text string argument by the key int argument to the left
func Decrypt(text string, key int) string {
	var result []rune
	runeArr := []rune(text)

	for _, r := range runeArr {
		// Check whether the rune r is in lowlist
		idx := indexOf(lowlist, r)
		if idx != -1 {
			if idx-key < 0 {
				result = append(result, lowlist[len(lowlist)-(key-idx)])
			} else {
				result = append(result, lowlist[idx-key])
			}
			continue
		}

		// Check whether the rune r is in capitallist
		idx = indexOf(capitallist, r)
		if idx != -1 {
			if idx-key < 0 {
				result = append(result, capitallist[len(capitallist)-(key-idx)])
			} else {
				result = append(result, capitallist[idx-key])
			}
			continue
		}

		// Save the rune r as it is
		result = append(result, r)
	}

	return string(result)
}

func indexOf(runeArr []rune, val rune) int {
	for i, r := range runeArr {
		if r == val {
			return i
		}
	}
	return -1
}
