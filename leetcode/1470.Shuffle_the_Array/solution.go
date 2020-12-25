package _470_Shuffle_the_Array

func shuffle(nums []int, n int) []int {
	result := make([]int, len(nums))

	i, j := 0,  n
	for k := 0; k < n; k += 1 {
		result[i], result[i+1] = nums[k], nums[j]
		j++
		i+=2
	}
	return  result
}
