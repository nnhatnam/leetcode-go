package _007_Reverse_Integer

func reverse(x int) int {
	var result int64 = 0
	for ; x != 0 ; {
		lastDigit := int64(x % 10)
		result = result * 10 + lastDigit
		if result > (2 << 30) - 1 || result < -(2 << 30){
			return 0
		}

		x = x / 10
	}
	return int(result)

}

