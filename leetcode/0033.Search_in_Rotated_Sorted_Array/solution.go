package _033_Search_in_Rotated_Sorted_Array

func binarySearch(nums []int, lo , hi , target int) int {
	//lo := 0
	//hi := len(nums)

	for ; lo <= hi ; {
		mid := int(uint(lo + hi ) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	lo := 0
	hi := len(nums) - 1

	for ; lo <= hi ; {
		mid := int(uint(lo + hi) >> 1)


		if nums[mid] == target {
			return mid
		} else if nums[lo] <= nums[mid] {
			//sorted array on the left side
			if nums[lo] <= target && target < nums[mid] {
				return binarySearch(nums, lo, mid - 1, target)
			}
			lo = mid + 1

		} else {
			//sorted array on the right side
			if nums[mid] < target && target <= nums[hi] {
				return binarySearch(nums, mid + 1, hi, target)
			}
			hi = mid - 1
		}
	}
	return -1
}
