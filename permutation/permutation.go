package permutation

// CheckPermutationAscii takes in two strings s1 and s2 as input parameter
// and tells us if s2 is a permutation of s1 or not
// NOTE: it is assumed that the strings only contain 8-byte printable characters
func CheckPermutationAscii(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count1 := make([]int, 256)
	count2 := make([]int, 256)

	for i := 0; i < len(s1); i++ {
		count1[s1[i]]++
		count2[s2[i]]++
	}

	for i := 0; i < 256; i++ {
		if count1[i] != count2[i] {
			return false
		}
	}

	return true
}
