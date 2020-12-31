package _588_Sum_of_All_Odd_Length_Subarrays

//[1,4,2,5,3]
func sumOddLengthSubarrays(arr []int) int {
	result := 0
	for i := 0 ; (2*i + 1) <= len(arr) ; i++ {

		sumCache := 0
		// [windowStart, windowEnd)
		for windowStart, windowEnd := 0, 2*i + 1 ; windowEnd <= len(arr); windowStart, windowEnd = windowStart + 1, windowEnd + 1 {
			if sumCache == 0 {
				for k := windowStart; k < windowEnd; k++ {
					sumCache += arr[k]
				}
			} else {
				sumCache = sumCache - arr[windowStart - 1] + arr[windowEnd - 1]
			}
			result += sumCache
		}
	}
	return result
}
