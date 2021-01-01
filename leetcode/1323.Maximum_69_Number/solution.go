package _323_Maximum_69_Number

func maximum69Number (num int) int {
	temp := []int{}
	for ; num > 0 ; num = num / 10 {
		temp = append(temp, num % 10)
	}
	result := 0
	isFlip := false
	for i := len(temp) - 1; i >= 0  ; i-- {
		if temp[i] == 6 && !isFlip {
			result = result*10 + 9
			isFlip = true
		} else {
			result = result*10 + temp[i]
		}
	}
	return result
}
