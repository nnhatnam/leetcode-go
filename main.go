package main

import "fmt"

//find how many number < v in array arr
func binarySearch(v int, arr []int) int{
	if len(arr) == 0 {
		return 0
	}

	result := 0
	lo := 0
	hi := len(arr) - 1
	mid := 0
	for ;lo <= hi; {
		mid = (lo + hi) / 2
		if arr[mid] < v {
			result = mid + 1
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	if l1 + l2 == 0 {
		return 0
	}
	if l2 > l1 {
		nums1, nums2 = nums2, nums1
	}

	mIndex := (l1 + l2 + 1) / 2 - 1

	lo := 0
	hi := mIndex
	r1 := mIndex
	for {
		v1 := nums1[r1]
		leftCount := binarySearch(v1, nums2)
		if r1 + leftCount == mIndex {

			if (l1 + l2) % 2 == 1 {
				return float64(nums1[r1])
			} else {
				if l1 == l2 {
					return float64(nums1[r1] + nums2[0]) / 2.0
				} else {
					return float64(nums1[r1] + nums1[r1 + 1]) / 2.0
				}
			}

		} else if r1 + leftCount > mIndex {
			hi = r1
			r1 = (lo + hi) / 2

		} else {

			lo = r1 + 1
			fmt.Println("lo ? ", lo, hi )
			r1 = (lo + hi) / 2
		}
	}

	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1 , 3} , []int{ 2 }))
}