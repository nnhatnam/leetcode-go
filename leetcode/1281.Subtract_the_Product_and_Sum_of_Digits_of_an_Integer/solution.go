package _281_Subtract_the_Product_and_Sum_of_Digits_of_an_Integer

func subtractProductAndSum(n int) int {
	sum := 0
	product := 1
	for ; n > 0 ; n = n / 10 {
		lastDigit := n % 10
		sum += lastDigit
		product *= lastDigit
	}
	return product - sum
}
