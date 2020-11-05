package _026_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	i := 0
	currentMax := nums[i]

	for i = 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			break
		}
		currentMax = nums[i]
	}

	for j := i ; j < len(nums); j++ {
		if nums[j] > currentMax {
			nums[i] , nums[j] = nums[j] , nums[i]
			currentMax = nums[i]
			i++
		}
	}

	return i
}
