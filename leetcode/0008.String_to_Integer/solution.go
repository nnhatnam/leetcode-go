package _008_String_to_Integer

import "strings"

func isOverFlow32(num int64) bool {
	return int64(num) > (2 << 30 -1) || int64(num) < -(2 << 30)
}

func myAtoi(s string) int {

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	numberStr := strings.Split(s, " ")[0]
	convertTable := map[string]int64{
		"0" : 0,
		"1" : 1,
		"2" : 2,
		"3" : 3,
		"4" : 4,
		"5" : 5,
		"6" : 6,
		"7" : 7,
		"8" : 8,
		"9" : 9,
	}

	var result int64 = 0

	sign := ""
	index := 0
	if s[0] == '-' {
		sign = "-"
		index = 1
	} else if s[0] == '+' {
		index = 1
	}

	for _, v := range numberStr[index:] {
		if num, ok := convertTable[string(v)]; ok {
			result = result*10 + num
			if isOverFlow32(result) {
				if sign == "-" {
					return -(2 << 30)
				}
				return (2 << 30) - 1
			}
		} else {
			break
		}
	}

	if sign == "-" {
		result = -result
		if isOverFlow32(result) {
			return -(2 << 30)
		}
	}

	return int(result)
}
