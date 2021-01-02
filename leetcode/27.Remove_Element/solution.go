package _7_Remove_Element

func removeElement(nums []int, val int) int {

	i, j := 0, len(nums) - 1

	for ; i <= j; i++ {

		if nums[i] == val {
			for ; nums[j] == val && j > i; j-- {}
			if j == i {
				return j
			}
			nums[i] , nums[j] = nums[j], nums[i]
			j--
		}

	}
	return j + 1

}
