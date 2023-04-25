package _014_Longest_Common_Prefix


func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		mIdx := mismatchIndex(prefix, strs[i])
		if mIdx == 0 {
			return ""
		}
		prefix = prefix[:mIdx]
	}

	return prefix
}

func mismatchIndex(s1 , s2 string) int {
	i := 0
	for ; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return i
}
