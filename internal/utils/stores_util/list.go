package stores_util

func RemoveDuplicates(intSlice []string) []string {
	seen := make(map[string]bool)
	var result []string

	for _, num := range intSlice {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}
