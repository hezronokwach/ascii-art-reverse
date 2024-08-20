package reverse

func CheckPattern(char, word []string) bool {
	if len(char[0]) > len(word[0]) {
		return false
	}
	for i, str := range word[:len(word)-1] {
		if char[i] != str[:len(char[i])] {
			return false
		}
	}
	return true
}

func TrimFound(length int, word []string) []string {
	for i, val := range word[0 : len(word)-1] {
		// fmt.Println(val[length:])
		word[i] = val[length:]
	}
	return word
}