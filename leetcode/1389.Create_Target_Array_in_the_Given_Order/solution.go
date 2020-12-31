package _389_Create_Target_Array_in_the_Given_Order

//[4,2,4,3,2] nums
//[0,0,1,3,1] index

func createTargetArray(nums []int, index []int) []int {

	for k, v := range index {
		if v < k {
			for j := 0 ; j < k ; j++ {
				if index[j] >= v {
					index[j]++
				}
			}
		}
	}

	target := make([]int, len(nums))

	for i, idx := range index  {
		target[idx] = nums[i]
	}

	return target
}


//Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
//func createTargetArray(nums []int, index []int) []int {
//	temp := make([]int, len(nums))
//	for i, idx := range index {
//		fmt.Println("id : ", i, idx)
//		if idx < i {
//			fmt.Println("ok ", idx)
//			for j := idx; j < i; j++ {
//				temp[j] += 1
//			}
//		}
//		temp[i] = idx //num at index i will be put to target at index idx
//		fmt.Println( " num at index " , i , " will be put to target at index ", idx,  temp)
//		//if idx before i, it will shift numbers from idx -> i - 1 to the right 1 step
//
//	}
//
//	fmt.Println("final: ", temp)
//	target := make([]int, len(nums))
//	for i, idx := range temp  {
//		target[idx] = nums[i]
//	}
//	return target
//}

//Not mine. It looks much more cleaner than my solution
//func createTargetArray(nums []int, index []int) []int {
//
//	arr := make([]int, 0)
//
//	for i0, i1 := range index {
//		arr = insert(arr, nums[i0], i1)
//	}
//	return arr
//}
//
//func insert(arr []int, val int, index int) []int {
//
//	newArr := make([]int, len(arr) + 1)
//	for i, v := range arr{
//		if i < index {
//			newArr[i] = v
//		} else {
//			newArr[i+1] = v
//		}
//	}
//	newArr[index] = val
//
//	return newArr
//
//}