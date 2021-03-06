package _4_Find_First_and_Last_Position_of_Element_in_Sorted_Array

//Not my solution. It's not worth it to do this problem
func searchRange(nums []int, target int) []int {
	nlen := len(nums)
	if nlen == 0 {
		return []int{-1, -1}
	}
	lo, hi := 0, nlen - 1
	var mid int


	for lo < hi && lo >= 0 && hi < nlen {
		if mid = lo + (hi - lo) / 2; nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < 0 || hi >= nlen || nums[lo] != target {
		return []int{-1, -1}
	}
	left := hi

	//search right
	lo, hi = left, nlen - 1
	for lo < hi && lo >= 0 && hi < nlen {
		if mid = hi - (hi - lo) / 2; nums[mid] <= target {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	right := lo

	return []int{left, right}
}