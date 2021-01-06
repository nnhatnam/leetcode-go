package _71_Sum_of_Two_Integers

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}

	return getSum(a ^ b, (a & b ) << 1)
}
