package reverse

func SliceFile(fileSlice []string) [][]string {
	var result [][]string
	for i := 0; i < len(fileSlice)-1; i += 9 {
		file := fileSlice[i+1 : i+9]
		result = append(result, file)
	}
	return result
}
