package _009_Palindrome_Number

func isPalindrome(x int) bool {

	if x < 0 || ( x != 0 && x % 10 == 0) {
		return false
	}

	reverseNum := 0

	for ; x > reverseNum; {
		lastDigit := x % 10
		reverseNum = reverseNum*10 + lastDigit
		x = x / 10
	}
	return x == reverseNum || x == reverseNum / 10
}