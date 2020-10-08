package Longest_Substring

//1abcabc
//01234

func lengthOfrunner(r int) int {
	return r + 1
}

func isInMap(key rune , m map[rune]int) bool {
	_ , ok := m[key]
	return ok
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	hashmap := map[rune]int{}
	r1 := 0
	r2 := 0
	max := 0


	for index , v := range s {
		if isInMap(v, hashmap) && hashmap[v] >= r1 {
			tempLength := lengthOfrunner(r2 - 1) - lengthOfrunner(r1) + 1
			if max < tempLength {
				max = tempLength
			}

			r1 = hashmap[v] + 1
		}

		hashmap[v] = index
		r2++
	}
	tempLength := lengthOfrunner(r2 - 1) - lengthOfrunner(r1) + 1
	if max < tempLength {
		max = tempLength
	}
	return max

}