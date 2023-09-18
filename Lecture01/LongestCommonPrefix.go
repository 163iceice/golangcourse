func longestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) < 1 {
		return result
	}

	for i, letter := range strs[0] {
		counter := 0
		for j := range strs {

			if j < (len(strs)-1) &&
				i < (len(strs[j+1])) &&
				(string(letter) == string(strs[j+1][i])) {
				counter++
			}
		}
		if counter == len(strs)-1 {
			result = fmt.Sprintf("%s%s", result, string(letter))
		} else {
			return result
		}
	}

	return result
}