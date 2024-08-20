package reverse

/*CheckPattern checks if a given character pattern matches the beginning of the words in the provided slice. It returns true if the pattern matches, and false otherwise.
The comparison is made line by line until a mismatch is found or all lines are checked.
*/
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