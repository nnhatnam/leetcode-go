package _013_Roman_to_Integer

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	hashmap := map[string]int {
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	sum := 0
	last_num := 0
	for i := len(s) - 1; i >= 0; i-- {
		num := hashmap[string(s[i])]

		if num < last_num {
			sum -= num
		} else {
			sum += num
		}
		last_num = num
	}
	return sum
}


