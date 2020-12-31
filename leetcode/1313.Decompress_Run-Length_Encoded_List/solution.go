package _313_Decompress_Run_Length_Encoded_List

func decompressRLElist(nums []int) []int {

	itemCount := 0
	//constraints: nums.length % 2 == 0
	for i := 0; i < len(nums) / 2 ; i++ {
		itemCount += nums[2*i]
	}
	result := make([]int, itemCount)
	currentPointer := 0
	for i := 0; i < len(nums) / 2 ; i++ {
		for j := 0; j < nums[2*i]; j++ {
			result[currentPointer] = nums[2*i + 1]
			currentPointer++
		}
	}
	return result
}