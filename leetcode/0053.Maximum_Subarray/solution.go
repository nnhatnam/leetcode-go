package _053_Maximum_Subarray

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	currentMax := nums[0]
	tempMax := nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax = max(nums[i] , nums[i] + tempMax)
		if tempMax > currentMax {
			currentMax = tempMax
		}
	}

	return currentMax

}