package _221_Split_a_String_in_Balanced_Strings

func balancedStringSplit(s string) int {
	count := 0
	balanced := 0
	for _, v := range s {
		if v == 'L' {
			balanced++
		} else {
			balanced--
		}
		if balanced == 0 {
			count++
		}
	}
	return count
}
