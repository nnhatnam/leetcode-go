package _017_Letter_Combinations_of_a_Phone_Number

func cartesian(s1, s2 []string) []string {
	result := []string{}
	for _, si := range s1 {
		for _, sj := range s2 {
			result = append(result, string(si) + string(sj))
		}
	}
	return result
}

func letterCombinations(digits string) []string {
	hashmap := map[string][]string{
		"2" : []string{"a" , "b" , "c"},
		"3" : []string{"d" , "e" , "f"},
		"4" : []string{"g" , "h" , "i"},
		"5" : []string{"j" , "k" , "l"},
		"6" : []string{"m" , "n" , "o"},
		"7" : []string{"p" , "q" , "r" , "s"},
		"8" : []string{"t" , "u" , "v"},
		"9" : []string{"w" , "x" , "y" , "z"},
	}
	if digits == "" {
		return []string{}
	}
	result := hashmap[string(digits[0])]
	for i := 1; i < len(digits); i++ {
		result = cartesian(result , hashmap[string(digits[i])])
	}
	return result
}
