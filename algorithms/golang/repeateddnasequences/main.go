package repeateddnasequences

func findRepeatedDnaSequences(s string) []string {
	var (
		record = make(map[string]int)
		subStr string
		result = make([]string, 0)
	)

	for i := 0; i+10 <= len(s); i++ {
		subStr = s[i : i+10]
		if cnt, ok := record[subStr]; ok && cnt == 1 {
			result = append(result, subStr)
		}

		record[subStr]++
	}

	return result
}
