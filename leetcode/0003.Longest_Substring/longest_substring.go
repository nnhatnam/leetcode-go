package Longest_Substring

//1abcabc
//01234

func lengthOfrunner(r int) int {
	return r + 1
}

func lengthBetweenTwoRunner(r1, r2 int) int {
	return r2 - r1 + 1
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

	for index, str := range s {

		dupIndex, ok := hashmap[str]

		if ok && dupIndex >= r1 {
			temp := lengthBetweenTwoRunner(r1, r2)
			if temp > max {
				max = temp
			}
			r1 = dupIndex + 1
		}
		hashmap[str] = index


		r2++
	}
	r2--
	temp := lengthBetweenTwoRunner(r1, r2)
	if temp > max {
		max = temp
	}

	return max

}