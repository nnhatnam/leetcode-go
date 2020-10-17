package _075_Sort_Colors

func sortColors(nums []int)  {
	l := len(nums)

	for i , j := 0 , 0 ; j < l;  {
		if nums[j] < 1 {
			nums[i] , nums[j] = nums[j] , nums[i]
			i++
			j++
		} else if nums[j] > 1 {
			l = l - 1
			nums[j] , nums[l] = nums[l], nums[j]
		} else {
			j++
		}
	}

}
